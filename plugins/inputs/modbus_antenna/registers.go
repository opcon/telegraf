package modbusAntenna

import (
	"bytes"
	"encoding/binary"
)

type register struct {
	addr   uint16
	label  string
	decode func(string, []byte) map[string]interface{}
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
func integer(name string, rawval []byte) map[string]interface{} {
	reader := bytes.NewReader(rawval)
	var val int32
	err := binary.Read(reader, binary.BigEndian, &val)
	if err != nil {
		//TODO: deal with error property
		panic(err)
	}
	return map[string]interface{}{name: val}
}
