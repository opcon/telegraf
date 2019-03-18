package modbusAntenna

import (
	"fmt"
	"time"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"

	"github.com/goburrow/modbus"
)

type modbusAntenna struct {
	AntennaType string
	Address     string
	Timeout     int
	MaxGap      int

	initDone bool
	groups   map[byte][][]register
}

var modbusAntennaConfig = `
  ## modbus antenna controller type
  antenna_type = "patriot12m"
  ## network address in form ip:port
  address = "192.168.1.22:502"
  ## modbus slave ID
  #slave_id = 0
  ##Timeout in milliseconds
  #timeout = 10000
  ## max gap between continuous regisers. Tweaking may improve performance
  #max_gap = 1
`

func (a *modbusAntenna) SampleConfig() string {
	return modbusAntennaConfig
}

func (a *modbusAntenna) Description() string {
	return "Query an antenna controller using modbus over TCP. Registers are assumed to be 32bits wide."
}

func (a *modbusAntenna) Gather(acc telegraf.Accumulator) error {
	var err error
	if !a.initDone {
		err = a.initAnt()
		if err != nil {
			return err
		}
	}

	fields := make(map[string]interface{})

	for slaveid, groups := range a.groups {
		handler := modbus.NewTCPClientHandler(a.Address)
		handler.SlaveId = slaveid
		handler.Timeout = time.Duration(a.Timeout) * time.Millisecond
		err = handler.Connect()
		if err != nil {
			return err
		}
		defer handler.Close()
		modbusClient := modbus.NewClient(handler)

		for _, group := range groups {
			startaddr := uint16(group[0].addr)
			endaddr := uint16(group[len(group)-1].addr)

			// Word = Size of register (32 bits)
			// ModbusWords = 16bits
			// For Patriot12m, the index number addresses ModbusWords, but you must
			// query a whole Word (32 bits) at a time. Reading half a word returns an error.

			const mbuswordsPerWord = 2
			numMbuswords := (endaddr - startaddr + 1) * mbuswordsPerWord

			// log.Println("I!: ", a.modbusClient, startaddr, numMbuswords)

			raw, err := modbusClient.ReadHoldingRegisters(startaddr, numMbuswords)
			if err != nil {
				return err
			}

			for _, register := range group {
				// Position in raw read
				// Do not assume the group is continuous
				const bytesPerWord = 4
				pos := (register.addr - startaddr) * bytesPerWord
				filtoutput := register.decode(register.label, raw[pos:pos+bytesPerWord])

				// Merge
				for k, v := range filtoutput {
					fields[k] = v
				}
			}
		}
	}
	acc.AddFields("antenna", fields, nil)
	return nil
}

func (a *modbusAntenna) initAnt() error {

	registers, ok := antennas[a.AntennaType]
	if !ok {
		return fmt.Errorf("unknown antenna %q", a.AntennaType)
	}

	a.groups = groupRegisters(registers, uint16(a.MaxGap))
	a.initDone = true
	return nil
}

func init() {
	inputs.Add("modbus_antenna", func() telegraf.Input { return &modbusAntenna{Timeout: 10000, MaxGap: 1} })
}
