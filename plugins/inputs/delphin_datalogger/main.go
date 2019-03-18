package datalogger

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"time"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/internal"
	"github.com/influxdata/telegraf/plugins/inputs"

	"github.com/goburrow/modbus"
)

// Standard MODBUS uses 16 bit words. Some are 32 bit though.
const bytesPerWord = 2

// Calibration from ISI 2018-10
var registers = map[string][]register{
	"temperature": {
		{0, "point1", fixedPoint(1e-2)},
		{1, "point2", fixedPoint(1e-2)},
		{2, "point3", fixedPoint(1e-2)},
		{3, "point4", fixedPoint(1e-2)},
		{4, "point5", fixedPoint(1e-2)},
		{5, "point6", fixedPoint(1e-2)},
		{6, "point7", fixedPoint(1e-2)},
		{7, "point8", fixedPoint(1e-2)},
		{8, "point9", fixedPoint(1e-2)},
		{9, "point10", fixedPoint(1e-2)},
		{10, "point11", fixedPoint(1e-2)},
		{11, "point12", fixedPoint(1e-2)},
	},

	"tilt": {
		{12, "x", fixedPoint(5.975e-4)},
		{13, "y", fixedPoint(6.025e-4)},
	},
}

type Datalogger struct {
	Address string
	Port    uint16
	Timeout internal.Duration
	SlaveId byte
}

var config = `
## Address and port of datalogger modbus port
address = "127.0.0.1"
port = 502
timeout = "20s"
slave_id = 1
`

func (s *Datalogger) SampleConfig() string {
	return config
}

func (s *Datalogger) Description() string {
	return "Query Delphin data logger configured from MGO"
}

func (s *Datalogger) Gather(acc telegraf.Accumulator) error {
	handler := modbus.NewTCPClientHandler(fmt.Sprintf("%s:%d", s.Address, s.Port))
	handler.SlaveId = s.SlaveId
	handler.Timeout = s.Timeout.Duration

	err := handler.Connect()
	if err != nil {
		return err
	}
	defer handler.Close()
	c := modbus.NewClient(handler)

	groupedFields := make(map[string]map[string]interface{})
	for name, group := range registers {
		startaddr := uint16(group[0].addr)
		endaddr := uint16(group[len(group)-1].addr)

		nwords := (endaddr - startaddr + 1)

		raw, err := c.ReadHoldingRegisters(startaddr, nwords)
		if err != nil {
			return err
		}

		fields := make(map[string]interface{})
		for _, register := range group {
			// Position in raw read
			// doesnt' assume the group is continuous
			pos := (register.addr - startaddr) * bytesPerWord
			fields[register.label], err = register.decode(raw[pos : pos+bytesPerWord])
			if err != nil {
				return err
			}
		}
		groupedFields[name] = fields
	}

	for name, fields := range groupedFields {
		acc.AddFields(fmt.Sprintf("antenna_%s", name), fields, nil)
	}

	return nil
}

type register struct {
	addr   uint16
	label  string
	decode func([]byte) (interface{}, error)
}

func fixedPoint(scale float64) func([]byte) (interface{}, error) {
	return func(rawval []byte) (interface{}, error) {
		var val int16
		reader := bytes.NewReader(rawval)
		err := binary.Read(reader, binary.BigEndian, &val)
		if err != nil {
			return nil, err
		}
		return float64(val) * scale, nil
	}
}

func init() {
	inputs.Add("delphin_datalogger", func() telegraf.Input {
		return &Datalogger{
			Address: "127.0.0.1",
			Port:    502,
			Timeout: internal.Duration{Duration: 20 * time.Second},
			SlaveId: 1,
		}
	})
}
