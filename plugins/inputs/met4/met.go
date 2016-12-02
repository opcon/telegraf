package met4

import (
	"bufio"
	"io"
	"net"
	"strconv"
	"strings"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
)

type Met4 struct {
	Address string
}

var Met4Config = `
  ## Address of metserver
  address = "127.0.0.1:50001"
  ## You can add extra tags, for example
  # [inputs.met4.tags]
  #     location = "..."
  #     device = "old"
`

func (s *Met4) SampleConfig() string {
	return Met4Config
}

var fieldNames []string = []string{
	"temperature",
	"pressure",
	"humidity",
	"windspeed",
	"windheading",
}

func (s *Met4) Description() string {
	return "Query at MET4 meteorological measurement systems via metserver"
}

func (s *Met4) Gather(acc telegraf.Accumulator) error {
	conn, err := net.Dial("tcp", s.Address)
	if err != nil {
		return err
	}

	line, err := bufio.NewReader(conn).ReadString('\n')
	if err == nil && err != io.EOF {
		return err
	}
	fields := parseLine(line)
	acc.AddFields("met", fields, nil)

	return nil
}

func init() {
	inputs.Add("met4", func() telegraf.Input { return &Met4{Address: "127.0.0.1:50001"} })
}

func parseLine(line string) map[string]interface{} {
	fields := make(map[string]interface{})
	if line == "" {
		return nil
	}

	fieldvals := strings.Split(line, ",")
	for i, f := range fieldvals {
		v, err := strconv.ParseFloat(f, 64)
		if err != nil {
			continue
		}
		if i >= len(fieldNames) {
			break
		}
		fields[fieldNames[i]] = v
	}
	return fields
}
