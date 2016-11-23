package fieldsystem

import (
	"log"
	"strings"
	"sync"
	"time"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/internal"
	"github.com/influxdata/telegraf/plugins/inputs"
)

type FieldSystem struct {
	Rdbe      bool
	Precision internal.Duration

	sync.Mutex
	wg   sync.WaitGroup
	done chan struct{}

	acc telegraf.Accumulator

	fs *Fscom
}

var FieldSystemConfig = `
  ## Poll the Field System state through shared memory. 
  ##
  ## Record RDBE Tsys and Pcal calculated by FS.
  ## This complements the rdbe_multicast input.
  #rdbe = false 
  ## Rate to poll FS variables.
  #precision = "100ms"
`

func (s *FieldSystem) SampleConfig() string {
	return FieldSystemConfig
}

func (s *FieldSystem) Description() string {
	return "Poll the Field System state through shared memory."
}

func (s *FieldSystem) Start(acc telegraf.Accumulator) (err error) {
	s.Lock()
	defer s.Unlock()

	s.done = make(chan struct{})

	s.acc = acc
	if s.Precision.Duration == 0 {
		s.Precision.Duration, _ = time.ParseDuration("100ms")
	}
	s.acc.SetPrecision(0, s.Precision.Duration)

	s.fs, err = GetSHM()
	if err != nil {
		return err
	}

	// FS semephores
	for i := 0; i < int(s.fs.Sem.Allocated); i++ {
		semname := strings.TrimSpace(string(s.fs.Sem.Name[i][:]))
		s.wg.Add(1)
		go s.watchsem(semname)
	}

	// FS bools
	fsbools := map[string]*int32{
		"data_valid": &s.fs.Data_valid[0].User_dv,
		"tracking":   &s.fs.Ionsor,
	}
	for name, b := range fsbools {
		s.wg.Add(1)
		go s.watchbool(name, b)
	}

	// FS strings
	fsstrings := map[string][]byte{
		"log":      s.fs.LLOG[:],
		"schedule": s.fs.LSKD[:],
		"source":   s.fs.Lsorna[:],
	}
	for name, str := range fsstrings {
		s.wg.Add(1)
		go s.watchstring(name, str)
	}

	if s.Rdbe {
		for i, _ := range s.fs.Rdbe_tsys_data {
			s.wg.Add(1)
			go s.watchrdbe(i)
		}
	}

	log.Println("Started FS listener service")
	return nil
}

func (s *FieldSystem) Stop() {
	s.Lock()
	defer s.Unlock()
	log.Println("Stopping FS listener service...")
	close(s.done)
	s.wg.Wait()
	log.Println("Stopped FS listener service")
}

func (s *FieldSystem) Gather(acc telegraf.Accumulator) (err error) {
	return nil
}

func init() {
	inputs.Add("fieldsystem", func() telegraf.Input { return &FieldSystem{} })
}
