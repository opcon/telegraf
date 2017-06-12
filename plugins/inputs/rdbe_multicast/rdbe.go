package rdbemulticast

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"math"
	"net"
	"sync"
	"time"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
)

// UDP_MAX_PACKET_SIZE UDP packet limit, see
// https://en.wikipedia.org/wiki/User_Datagram_Protocol#Packet_structure
const UDP_MAX_PACKET_SIZE int = 64 * 1024

const sampleConfig = `
  ## RDBE devices to listen. Can be an ID or a multicast address and port
  # eg.
  # device_ids = ["a","b","c","d"]
  # device_ids = ["239.0.2.40:20024"]
  device_ids = ["a","b","c","d"]
  ## Save Tsys, Pcal, and Raw measurments
  ## these are saved into the "rdbe_multicast_*" measurment
  save_pcal = false
  save_tsys = false
  save_raw = false
  save_statstr = false

  ## Extra tags should be added
  ## eg.
  #[inputs.rdbe.tags]
  #  antenna = "gs"
  #  foo = "bar"
`

// layout of RDBE packet
type rdbepacket struct {
	ReadTime      [20]byte
	PacketSize    uint16
	EpochRef      uint16
	EpochSec      uint32
	Interval      uint32
	TsysHeader    [20]byte
	TsysOn        [64]int32
	TsysOff       [64]int32
	PcalHeader    [20]byte
	PcalIfx       uint16
	PcalIfxPad    uint16
	PcalSin       [1024]int32
	PcalCos       [1024]int32
	StatStr       [3000]byte
	RawHeader     [24]byte
	RawIfx        uint16
	RawIfxPad     uint16
	Mu            float64
	Sigma         float64
	PpsOffset     float64
	GpsOffset     float64
	RawSize       uint16
	RawSamples    [4096]byte
	RawSamplesPad [6]byte
}

// Get the multicast address from the device ID
func broadcastAddress(id string) (*net.UDPAddr, error) {
	// Try lookup first
	addr, err := net.ResolveUDPAddr("udp", id)
	if err == nil {
		return addr, nil
	}

	bid := byte(id[0])
	if bid < 'a' || bid > 'z' {
		return nil, errors.New("Invalid multicast address")
	}
	ip := fmt.Sprintf("239.0.2.%d", (bid-'a'+1)*10)
	port := 20021 + int(bid-'a')
	addr, err = net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", ip, port))
	return addr, err

}

func cstr(str []byte) string {
	var n int
	for n = 0; str[n] != 0; n++ {
	}
	if n == 0 {
		return ""
	}
	return string(str[:n-1])
}

// Convert a slice to a map, with the keys as zero
// padded decimal strings.
//
// Needed for sorting in InfluxDB
func intslicetomap(s []int32) map[string]interface{} {
	m := map[string]interface{}{}

	// Number of digits required to display list
	slen := int(math.Floor(math.Log10(float64(len(s))))) + 1
	fmtstr := fmt.Sprintf("%%0%dd", slen)
	for i := range s {
		m[fmt.Sprintf(fmtstr, i)] = s[i]
	}
	return m
}

func byteslicetomap(s []byte) map[string]interface{} {
	m := map[string]interface{}{}

	// Number of digits required to display list
	slen := int(math.Floor(math.Log10(float64(len(s))))) + 1
	fmtstr := fmt.Sprintf("%%0%dd", slen)
	for i := range s {
		m[fmt.Sprintf(fmtstr, i)] = s[i]
	}
	return m
}

func copymap(f map[string]string) map[string]string {
	t := map[string]string{}
	for k, v := range f {
		t[k] = v
	}
	return t
}

// RdbeMulticast Based on UDP listener
type RdbeMulticast struct {
	DeviceIds   []string
	SaveRaw     bool
	SavePcal    bool
	SaveTsys    bool
	SaveStatstr bool

	in   chan []byte
	done chan struct{}

	listeners map[string]*net.UDPConn

	// Keep the accumulator in this struct
	telegraf.Accumulator

	sync.Mutex
	wg sync.WaitGroup
}

func (u *RdbeMulticast) SampleConfig() string {
	return sampleConfig
}

func (u *RdbeMulticast) Description() string {
	return "RDBE UDP Multicast listener"
}

// All the work is done in the Start() function, so this is just a dummy to satisfy the interface
func (u *RdbeMulticast) Gather(_ telegraf.Accumulator) error {
	return nil
}

func (u *RdbeMulticast) Start(acc telegraf.Accumulator) error {
	u.Lock()
	defer u.Unlock()

	u.Accumulator = acc
	u.done = make(chan struct{})
	u.listeners = make(map[string]*net.UDPConn)

	for _, id := range u.DeviceIds {
		u.wg.Add(1)
		go u.rdbeListen(id)
	}

	log.Println("I! Started RDBE Multicast listener service")
	return nil
}

func (u *RdbeMulticast) Stop() {
	u.Lock()
	defer u.Unlock()
	close(u.done)
	u.wg.Wait()
	log.Println("I! Stopped RDBE Multicast listener service")
}

func (u *RdbeMulticast) rdbeListen(id string) {
	defer u.wg.Done()

	addr, err := broadcastAddress(id)
	if err != nil {
		u.AddError(fmt.Errorf("unable to construct address from \"%s\": %s", id, err))
		return
	}

	var conn *net.UDPConn
conloop:
	for {
		select {
		case <-u.done:
			return
		case <-time.Tick(60 * time.Second):
			conn, err := net.ListenMulticastUDP("udp", nil, addr)
			if err != nil {
				u.AddError(fmt.Errorf("unable to start UDP listener: %s", err))
				continue
			}
			break
		}
	}
	defer conn.Close()

	log.Printf("I! RDBE Multicast listening to %s\n", addr.String())

	buf := make([]byte, UDP_MAX_PACKET_SIZE)
	pack := rdbepacket{}

	for {
		select {
		case <-u.done:
			return
		default:
			n, err := conn.Read(buf)
			if err != nil {
				u.AddError(fmt.Errorf("unable to read from conn: %s", err))
				continue
			}
			now := time.Now()

			reader := bytes.NewReader(buf)
			err = binary.Read(reader, binary.BigEndian, &pack)
			if err != nil {
				u.AddError(fmt.Errorf("unable to unpack RDBE packet: %s", err))
				continue
			}

			if n != int(pack.PacketSize) {
				u.AddError(fmt.Errorf("bad RDBE packet length"))
				continue
			}

			tags := map[string]string{
				"id":      id,
				"pcalifx": fmt.Sprintf("%d", pack.PcalIfx),
				"rawifx":  fmt.Sprintf("%d", pack.RawIfx),
			}

			fields := map[string]interface{}{
				"readtime":  cstr(pack.ReadTime[:]),
				"epochref":  pack.EpochRef,
				"epochsec":  pack.EpochSec,
				"interval":  pack.Interval,
				"mu":        pack.Mu,
				"sigma":     pack.Sigma,
				"ppsoffset": pack.PpsOffset,
				"gpsoffset": pack.GpsOffset,
			}
			u.AddFields("rdbe_multicast", fields, tags, now)

			if u.SaveRaw {
				rawfields := byteslicetomap(pack.RawSamples[:])
				u.AddFields("rdbe_multicast_raw", rawfields, tags, now)
			}

			if u.SavePcal {
				pcalcos := intslicetomap(pack.PcalCos[:])
				t := copymap(tags)
				t["component"] = "cos"
				u.AddFields("rdbe_multicast_pcal", pcalcos, t, now)

				pcalsin := intslicetomap(pack.PcalSin[:])
				t = copymap(tags)
				t["component"] = "sin"
				u.AddFields("rdbe_multicast_pcal", pcalsin, t, now)
			}

			if u.SaveTsys {
				fieldsOn := intslicetomap(pack.TsysOn[:])
				t := copymap(tags)
				t["cal"] = "on"
				u.AddFields("rdbe_multicast_tsys", fieldsOn, t, now)

				fieldsOff := intslicetomap(pack.TsysOff[:])
				t = copymap(tags)
				t["cal"] = "off"
				u.AddFields("rdbe_multicast_tsys", fieldsOff, t, now)

			}
			if u.SaveStatstr {
				rawfields := map[string]interface{}{
					"statstr": cstr(pack.StatStr[:]),
				}
				u.AddFields("rdbe_multicast_statstr", rawfields, tags, now)
			}
		}
	}
}

func init() {
	inputs.Add("rdbe_multicast", func() telegraf.Input {
		return &RdbeMulticast{}
	})
}
