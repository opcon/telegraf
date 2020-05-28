package dbbc3multicast

import (
	"encoding/json"
	"io"
	"log"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
	jsonparser "github.com/influxdata/telegraf/plugins/parsers/json"

	"github.com/dehorsley/dbbc3mcast"
)

// UDP_MAX_PACKET_SIZE UDP packet limit, see
// https://en.wikipedia.org/wiki/User_Datagram_Protocol#Packet_structure
const UDP_MAX_PACKET_SIZE int = 64 * 1024

const sampleConfig = `
  # multicast group and port to listen on. 
  address = "224.0.0.255:25000"
  ## GJSON query to manipulate the logged data
  # json_query = ""

`

type DbbcMulticast struct {
	Address string `toml:"address"`
	Query   string `toml:"json_query"`
	l       *dbbc3mcast.Dbbc3DDCMulticastListener
	telegraf.Accumulator

	Log telegraf.Logger
}

func (u *DbbcMulticast) SampleConfig() string {
	return sampleConfig
}

func (u *DbbcMulticast) Description() string {
	return "DBBC3 Multicast listener"
}

// All the work is done in the Start() function, so this is just a dummy to satisfy the interface
func (u *DbbcMulticast) Gather(_ telegraf.Accumulator) error {
	return nil
}

func (u *DbbcMulticast) Start(acc telegraf.Accumulator) error {
	c := &jsonparser.Config{
		MetricName:   "dbbc3multicast",
		TagKeys:      []string{"version"},
		NameKey:      "",
		StringFields: []string{},
		Query:        u.Query,
		TimeKey:      "",
		TimeFormat:   "",
		Timezone:     "",
		DefaultTags:  map[string]string{},
		Strict:       false,
	}
	parser, err := jsonparser.New(c)
	if err != nil {
		return err
	}

	l, err := dbbc3mcast.New(u.Address)
	if err != nil {
		u.Log.Errorf("starting DBBC3 multicast listener %w", err)
		return err
	}
	u.l = l
	go func() {
		for {
			msg, err := u.l.Next()
			if err == io.EOF {
				return
			}
			if err != nil {
				u.Log.Errorf("reading DBBC3 multicast: %w", err)
				continue
			}

			msgjson, err := json.Marshal(&msg)

			ms, err := parser.Parse(msgjson)
			if err != nil {
				u.Log.Errorf("Error parsing json: %w", err)
				continue
			}

			if ms == nil || len(ms) == 0 {
				continue
			}
			for _, m := range ms {
				acc.AddFields(m.Name(), m.Fields(), m.Tags(), m.Time())
			}
		}
	}()
	return nil
}

func (u *DbbcMulticast) Stop() {
	log.Println("I! Stopped DBBC3 Multicast listener service")
}

func init() {
	inputs.Add("dbbc3_multicast", func() telegraf.Input {
		return &DbbcMulticast{}
	})
}
