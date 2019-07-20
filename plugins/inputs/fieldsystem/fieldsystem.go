package fieldsystem

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/internal"
	"github.com/influxdata/telegraf/plugins/inputs"

	"github.com/pkg/errors"

	fs "github.com/nvi-inc/fsgo"
)

type FieldSystem struct {
	Rdbe      bool
	Precision internal.Duration
	Version   string

	sync.Mutex
	done chan struct{}
	wg   sync.WaitGroup

	acc telegraf.Accumulator

	fs fs.FieldSystem
}

var FieldSystemConfig = `
## Poll the Field System state through shared memory. 
##
## Record RDBE Tsys and Pcal calculated by FS.
## This complements the rdbe_multicast input.
#rdbe = false 
## Rate to poll FS variables.
#precision = "100ms"
#version = "9.13.0"
`

func (s *FieldSystem) SampleConfig() string {
	return FieldSystemConfig
}

func (s *FieldSystem) Description() string {
	return "Poll the Field System state through shared memory."
}

const fsMeas = "fs"

func (s *FieldSystem) Start(acc telegraf.Accumulator) (err error) {
	s.Lock()
	defer s.Unlock()

	s.done = make(chan struct{})

	s.acc = acc
	s.acc.SetPrecision(s.Precision.Duration)

	if s.Version == "" {
		s.fs, err = fs.NewFieldSystem()

		if err != nil {
			return errors.Wrap(err, "connecting to the field system")
		}
	} else {
		s.fs, err = fs.NewFieldSystemVersion(s.Version)

		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("connecting to the field system with version %q", s.Version))
		}
	}

	// FS semephores

	sems := s.fs.Semaphores()
	for i := range sems {
		semname := sems[i]
		s.watch(fsMeas, semname, func() (interface{}, error) {
			return s.fs.SemLocked(semname)
		})
	}

	// FS bools
	s.watch(fsMeas, "data_valid", func() (interface{}, error) { return s.fs.DataValid(), nil })
	s.watch(fsMeas, "tracking", func() (interface{}, error) { return s.fs.Tracking(), nil })

	// FS strings
	s.watch(fsMeas, "log", func() (interface{}, error) { return s.fs.Log(), nil })
	s.watch(fsMeas, "schedule", func() (interface{}, error) { return s.fs.Schedule(), nil })
	s.watch(fsMeas, "source", func() (interface{}, error) { return s.fs.Source(), nil })

	if r, ok := s.fs.(fs.Rdbe); ok {
		if s.Rdbe {
			for i := 0; i < r.RdbeNum(); i++ {
				s.watchrdbe(i)
			}
		}
	}

	log.Println("I! Started FS listener service...")
	return nil
}

func (s *FieldSystem) watch(meas string, name string, get func() (interface{}, error)) {
	s.wg.Add(1)
	go func() {
		var oldval interface{}
		fields := map[string]interface{}{}
		defer s.wg.Done()
		for {
			select {
			case <-s.done:
				return
			case <-time.After(s.Precision.Duration):
				val, err := get()
				if err != nil {
					log.Println(err)
				}
				if val != oldval {
					fields[name] = val
					s.acc.AddFields(meas, fields, nil, time.Now())
					oldval = val
				}
			}
		}
	}()
}

const (
	rdbeMeas          = "fs_rdbe"
	tsysRdbeMeas      = "fs_rdbe_tsys"
	pcalPhaseRdbeMeas = "fs_rdbe_pcal_phase"
	pcalAmpRdbeMeas   = "fs_rdbe_pcal_amp"
)

func (s *FieldSystem) watchrdbe(rdbeindex int) {
	rdbeid := string('a' + rdbeindex)

	r := s.fs.(fs.Rdbe)

	updated, err := r.RdbeUpdatedFn(rdbeindex)
	if err != nil {
		panic(err)
	}

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		for {
			select {
			case <-s.done:
				return
			case <-time.After(s.Precision.Duration):
				if !updated() {
					continue
				}

				fields := map[string]interface{}{}

				tags := map[string]string{
					"rdbe": rdbeid,
				}

				rdbedata, err := r.RdbeMap(rdbeindex)
				if err != nil {
					log.Println("E! ", err)
				}

				now := time.Now()

				fields["epoch"] = rdbedata["Epoch"]
				fields["epoch_vdif"] = rdbedata["EpochVdif"]
				fields["pcaloff"] = rdbedata["Pcaloff"]
				fields["dot2gps"] = rdbedata["Dot2gps"]
				fields["sigma"] = rdbedata["Sigma"]
				//fields["pcal_spacing"] = rdbedata.Pcal_spacing

				// Epoch      [14]byte
				// pad_cgo_0  [2]byte
				// Epoch_vdif int32
				// Tsys       [18][2]float32
				// Pcal_amp   [1024]float32
				// Pcal_phase [1024]float32
				// Pcal_ifx   int32
				// Sigma      float32
				// Raw_ifx    int32
				// Dot2gps    float64
				// Pcaloff    float64

				// Tsys data
				// Both IFs are updated every seconds
				tsys := rdbedata["Tsys"].([][]float32)
				var tsysfields map[string]interface{}
				for i := 0; i < len(tsys[0]); i++ {
					tsysfields = map[string]interface{}{}
					tags = map[string]string{
						"rdbe": rdbeid,
						"if":   fmt.Sprintf("%d", i),
					}
					for j := 0; j < len(tsys); j++ {
						if tsys[j][i] < -8e6 || IsNan32(tsys[j][i]) || IsInf32(tsys[j][i]) {
							continue
						}
						tsysfields[fmt.Sprintf("chan_%04d", j)] = tsys[j][i]
					}
					s.acc.AddFields(tsysRdbeMeas, tsysfields, tags, now)
				}

				// Pcal data
				// alternates between IF every second
				pcalfields := map[string]interface{}{}
				tags = map[string]string{
					"rdbe": rdbeid,
					"if":   fmt.Sprintf("%d", rdbedata["PcalIfx"]),
				}

				pcalamp := rdbedata["PcalAmp"].([]float32)
				for i, amp := range pcalamp {
					if amp < -8e6 || IsNan32(amp) || IsInf32(amp) {
						delete(pcalfields, fmt.Sprintf("chan_%04d", i))
						continue
					}
					pcalfields[fmt.Sprintf("chan_%04d", i)] = amp
				}
				s.acc.AddFields(pcalAmpRdbeMeas, pcalfields, tags, now)

				pcalphase := rdbedata["PcalPhase"].([]float32)
				for i, phase := range pcalphase {
					if phase < -8e6 || IsNan32(phase) || IsInf32(phase) {
						delete(pcalfields, fmt.Sprintf("chan_%04d", i))
						continue
					}
					pcalfields[fmt.Sprintf("chan_%04d", i)] = phase
				}
				s.acc.AddFields(pcalPhaseRdbeMeas, pcalfields, tags, now)

			}
		}
	}()
}

func (s *FieldSystem) Stop() {
	s.Lock()
	defer s.Unlock()
	log.Println("I! Stopping FS listener service...")
	close(s.done)
	s.wg.Wait()
	log.Println("I! Stopped FS listener service.")
}

func (s *FieldSystem) Gather(acc telegraf.Accumulator) (err error) {
	return nil
}

func init() {
	inputs.Add("fieldsystem", func() telegraf.Input {
		p, _ := time.ParseDuration("100ms")
		return &FieldSystem{Precision: internal.Duration{p}}
	})
}
