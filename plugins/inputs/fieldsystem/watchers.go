package fieldsystem

import (
	"fmt"
	"log"
	"time"
)

const fsMeas = "fs"

func (s *FieldSystem) watch(meas string, name string, get func() (interface{}, error)) {
	defer s.wg.Done()
	fields := map[string]interface{}{}
	var oldval interface{}
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
}

// Watch a string in the FS shm struct
func (s *FieldSystem) watchstring(name string, str []byte) {
	s.watch(fsMeas, name, func() (interface{}, error) { return fsstr(str), nil })
}

func (s *FieldSystem) watchsem(name string) {
	s.watch(fsMeas, name, func() (interface{}, error) { return s.fs.SemLocked(name) })

}

func (s *FieldSystem) watchbool(name string, addr *int32) {
	s.watch(fsMeas, name, func() (interface{}, error) { return (*addr) == 1, nil })
}

func (s *FieldSystem) watchint(name string, addr *int32) {
	s.watch(fsMeas, name, func() (interface{}, error) { return *addr, nil })
}

const (
	rdbeMeas          = "fs_rdbe"
	tsysRdbeMeas      = "fs_rdbe_tsys"
	pcalPhaseRdbeMeas = "fs_rdbe_pcal_phase"
	pcalAmpRdbeMeas   = "fs_rdbe_pcal_amp"
)

func (s *FieldSystem) watchrdbe(rdbeindex int) {
	defer s.wg.Done()
	rdbeid := string('a' + rdbeindex)

	oldiping := int32(-1)

	for {
		select {
		case <-s.done:
			return
		case <-time.After(s.Precision.Duration):
			iping := s.fs.Rdbe_tsys_data[rdbeindex].Iping
			if iping == oldiping {
				continue
			}

			oldiping = iping
			if iping < 0 {
				continue
			}
			fields := map[string]interface{}{}

			tags := map[string]string{
				"rdbe": rdbeid,
			}

			rdbedata := s.fs.Rdbe_tsys_data[rdbeindex].Data[iping]
			now := time.Now()

			fields["epoch"] = cstr(rdbedata.Epoch[:])
			fields["epoch_vdif"] = rdbedata.Epoch_vdif
			fields["pcaloff"] = rdbedata.Pcaloff
			fields["dot2gps"] = rdbedata.Dot2gps
			fields["sigma"] = rdbedata.Sigma
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
			var tsysfields map[string]interface{}
			for i := 0; i < len(rdbedata.Tsys[0]); i++ {
				tsysfields = map[string]interface{}{}
				tags = map[string]string{
					"rdbe": rdbeid,
					"if":   fmt.Sprintf("%d", i),
				}
				for j := 0; j < len(rdbedata.Tsys); j++ {
					if rdbedata.Tsys[j][i] < -8e6 || IsNan32(rdbedata.Tsys[j][i]) || IsInf32(rdbedata.Tsys[j][i]) {
						continue
					}
					tsysfields[fmt.Sprintf("chan_%04d", j)] = rdbedata.Tsys[j][i]
				}
				s.acc.AddFields(tsysRdbeMeas, tsysfields, tags, now)
			}

			// Pcal data
			// alternates between IF every second
			pcalfields := map[string]interface{}{}
			tags = map[string]string{
				"rdbe": rdbeid,
				"if":   fmt.Sprintf("%d", rdbedata.Pcal_ifx),
			}
			for i, amp := range rdbedata.Pcal_amp {
				if amp < -8e6 || IsNan32(amp) || IsInf32(amp) {
					delete(pcalfields, fmt.Sprintf("chan_%04d", i))
					continue
				}
				pcalfields[fmt.Sprintf("chan_%04d", i)] = amp
			}

			s.acc.AddFields(pcalAmpRdbeMeas, pcalfields, tags, now)

			for i, phase := range rdbedata.Pcal_phase {
				if phase < -8e6 || IsNan32(phase) || IsInf32(phase) {
					delete(pcalfields, fmt.Sprintf("chan_%04d", i))
					continue
				}
				pcalfields[fmt.Sprintf("chan_%04d", i)] = phase
			}
			s.acc.AddFields(pcalPhaseRdbeMeas, pcalfields, tags, now)

		}
	}
}
