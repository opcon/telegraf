package modbusAntenna

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type decodeFn func(string, []byte) (map[string]interface{}, error)

// currently just storing by description
type register struct {
	addr        uint16
	id          string
	description string
	decode      decodeFn
}

// Group continuous registers together
// Registers are assumed to be sorted.
func groupRegisters(slaves map[byte][]register, maxgap uint16) map[byte][][]register {
	slaveGroups := make(map[byte][][]register)

	for slaveID, registers := range slaves {
		expectAddr := uint16(0)
		groups := make([][]register, 0, len(registers))
		group := make([]register, 0, len(registers))

		for i, reg := range registers {
			if i == 0 || reg.addr-expectAddr <= maxgap {
				group = append(group, reg)
			} else {
				groups = append(groups, group)
				group = make([]register, 0, len(registers))
				group = append(group, reg)
			}
			expectAddr = reg.addr + 1
		}
		groups = append(groups, group)
		slaveGroups[slaveID] = groups
	}
	return slaveGroups

}

// Filter functions

// Interpret registers as a big-endian 32bit int
func integer(name string, rawval []byte) (map[string]interface{}, error) {
	reader := bytes.NewReader(rawval)
	var val int32
	err := binary.Read(reader, binary.BigEndian, &val)
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{name: val}, nil
}

func boolean(name string, rawval []byte) (map[string]interface{}, error) {
	reader := bytes.NewReader(rawval)
	var val uint32
	err := binary.Read(reader, binary.BigEndian, &val)
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{name: val != 0}, nil
}

// Interpret registers as a real values stored as a fixed-point number,
// encoded in a big-endian 32bit with a scaling factor given by scale
func fixedpoint(scale float64) decodeFn {
	return func(name string, rawval []byte) (map[string]interface{}, error) {
		var val int32
		reader := bytes.NewReader(rawval)
		err := binary.Read(reader, binary.BigEndian, &val)
		if err != nil {
			return nil, err
		}
		return map[string]interface{}{name: float64(val) * scale}, nil
	}
}

func enum(names ...string) decodeFn {
	return func(name string, rawval []byte) (map[string]interface{}, error) {
		reader := bytes.NewReader(rawval)
		var val uint32
		err := binary.Read(reader, binary.BigEndian, &val)
		if err != nil {
			return nil, err
		}

		if val >= uint32(len(names)) {
			return nil, fmt.Errorf("out of range")
		}

		return map[string]interface{}{name: names[val]}, nil
	}
}
