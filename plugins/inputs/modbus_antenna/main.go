package modbus_antenna

import (
	"errors"
	"time"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"

	"github.com/goburrow/modbus"
)

type ModbusAntenna struct {
	AntennaType string
	Address     string
	SlaveId     int
	Timeout     int
	MaxGap      int

	initDone     bool
	groups       [][]register
	modbusClient modbus.Client
}

var ModbusAntennaConfig = `
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

func (a *ModbusAntenna) SampleConfig() string {
	return ModbusAntennaConfig
}

func (a *ModbusAntenna) Description() string {
	return "Query an antenna controller using modbus over TCP. Registers are assumed to be 32bits wide."
}

func (a *ModbusAntenna) Gather(acc telegraf.Accumulator) error {
	var err error
	if a.initDone == false {
		err = a.initAnt()
		if err != nil {
			return err
		}
	}

	fields := make(map[string]interface{})

	for _, group := range a.groups {
		startaddr := uint16(group[0].addr)
		endaddr := uint16(group[len(group)-1].addr)

		// Word = Size of register (32 bits)
		// ModbusWords = 16bits
		// For Patriot12m, the index number addresses ModbusWords, but you must
		// query a whole Word (32 bits) at a time. Reading half a word returns an error.

		const mbuswordsPerWord = 2
		numMbuswords := (endaddr - startaddr + 1) * mbuswordsPerWord

		raw, err := a.modbusClient.ReadHoldingRegisters(startaddr, numMbuswords)
		if err != nil {
			err2 := a.initConn()
			if err2 != nil {
				return err2
			}
			return err
		}

		for _, register := range group {
			// Position in raw read
			// Do not assume the group is continuous
			const bytesPerWord = 4
			pos := (register.addr - startaddr) * bytesPerWord
			filtoutput := register.filter(register.label, raw[pos:pos+bytesPerWord])

			// Merge
			for k, v := range filtoutput {
				fields[k] = v
			}
		}
	}
	acc.AddFields("antenna", fields, nil)
	return nil
}

func (a *ModbusAntenna) initAnt() (err error) {

	registers, ok := antennas[a.AntennaType]
	if !ok {
		return errors.New("unknown antenna")
	}

	a.groups, err = groupRegisters(registers, uint16(a.MaxGap))
	if err != nil {
		return
	}

	err = a.initConn()

	a.initDone = true
	return
}

func (a *ModbusAntenna) initConn() (err error) {
	handler := modbus.NewTCPClientHandler(a.Address)
	handler.SlaveId = byte(a.SlaveId)

	handler.Timeout = time.Duration(a.Timeout) * time.Millisecond
	err = handler.Connect()
	if err != nil {
		return
	}
	a.modbusClient = modbus.NewClient(handler)
	return
}

func init() {
	inputs.Add("modbus_antenna", func() telegraf.Input { return &ModbusAntenna{Timeout: 10000, MaxGap: 1} })
}
