package fieldsystem

import (
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
)

type FieldSystem struct {
}

var FieldSystemConfig = `
  ##Stuff
`

func (s *FieldSystem) SampleConfig() string {
	return FieldSystemConfig
}

func (s *FieldSystem) Description() string {
	return ""
}

func (s *FieldSystem) Gather(acc telegraf.Accumulator) error {
	fields := make(map[string]interface{})
	tags := make(map[string]string)

	fs, err := GetFSSHM()
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", fs.LLOG[:])
	for i := 0; i < int(fs.Sem.Allocated); i++ {
		v := fs.Sem.Name[i]
		i, _ := fs.SemLocked(string(v[:]))
		fmt.Printf("%s, %t\n", v[:], i)
	}

	acc.AddFields("fs", fields, tags)
	return nil
}

func init() {
	inputs.Add("fieldsystem", func() telegraf.Input { return &FieldSystem{} })
}
