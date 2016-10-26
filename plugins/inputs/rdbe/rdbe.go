package rdbe

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
	"github.com/influxdata/telegraf/plugins/parsers"
)

// Rdbe Based on UDP listener
type Rdbe struct {
	DeviceIds              []string
	AllowedPendingMessages int
	SaveRaw                bool

	sync.Mutex
	wg sync.WaitGroup

	in   chan []byte
	done chan struct{}
	// drops tracks the number of dropped metrics.
	drops int
	// malformed tracks the number of malformed packets
	malformed int

	parser parsers.Parser

	// Keep the accumulator in this struct
	acc telegraf.Accumulator

	listeners map[string]*net.UDPConn
}

// UDP_MAX_PACKET_SIZE UDP packet limit, see
// https://en.wikipedia.org/wiki/User_Datagram_Protocol#Packet_structure
const UDP_MAX_PACKET_SIZE int = 64 * 1024

const sampleConfig = `
  ## IDs of RDBE devices to listen 
  device_ids = ["a","b","c","d"]
  ## Save raw Tsys and Pcal measurments
  ## these are saved into the "rdbe_raw" measurment
  save_raw = false

  ## Extra tags should be added
  ## eg.
  #[inputs.rdbe.tags]
  #  antenna = "gs"
  #  foo = "bar"
`

func (u *Rdbe) SampleConfig() string {
	return sampleConfig
}

func (u *Rdbe) Description() string {
	return "RDBE UDP Multicast listener"
}

// All the work is done in the Start() function, so this is just a dummy
// function.
func (u *Rdbe) Gather(_ telegraf.Accumulator) error {
	return nil
}

func (u *Rdbe) SetParser(parser parsers.Parser) {
	u.parser = parser
}

func (u *Rdbe) Start(acc telegraf.Accumulator) error {
	u.Lock()
	defer u.Unlock()

	u.acc = acc
	u.in = make(chan []byte, u.AllowedPendingMessages)
	u.done = make(chan struct{})
	u.listeners = make(map[string]*net.UDPConn)

	for _, id := range u.DeviceIds {
		u.wg.Add(1)
		go u.rdbeListen(id)
	}

	log.Println("Started RDBE listener service")
	return nil
}

func (u *Rdbe) Stop() {
	u.Lock()
	defer u.Unlock()
	close(u.done)
	u.wg.Wait()
	close(u.in)
	log.Println("Stopped RDBE listener service")
}

type rdbepacket struct {
	ReadTime      [20]byte
	PacketSize    uint16
	EpochRef      uint16
	EpochSec      uint32
	Interval      uint32
	TsysHeader    [20]byte
	TsysOn        [64]uint32
	TsysOff       [64]uint32
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
func broadcastAddress(id string) string {
	bid := byte(id[0])
	if bid < 'a' || bid > 'z' {
		log.Fatal("bad rdbe id %s", id)
	}
	addr := fmt.Sprintf("239.0.2.%d", (bid-'a'+1)*10)
	port := 20021 + int(bid-'a')
	return fmt.Sprintf("%s:%d", addr, port)
}

func (u *Rdbe) rdbeListen(id string) {
	defer u.wg.Done()

	addr, _ := net.ResolveUDPAddr("udp", broadcastAddress(id))
	conn, err := net.ListenMulticastUDP("udp", nil, addr)
	defer conn.Close()
	if err != nil {
		log.Fatalf("ERROR: ListenUDP - %s", err)
	}
	log.Printf("RDBE server listening to %s", conn.LocalAddr().String())

	var pack rdbepacket
	buf := make([]byte, UDP_MAX_PACKET_SIZE)

	for {
		select {
		case <-u.done:
			return
		default:
			n, err := conn.Read(buf)
			if err != nil {
				log.Printf("ERROR: %s\n", err.Error())
				continue
			}

			reader := bytes.NewReader(buf)
			err = binary.Read(reader, binary.BigEndian, &pack)
			if err != nil {
				log.Printf("ERROR: %s\n", err)
				continue
			}

			if n != int(pack.PacketSize) {
				log.Println("ERROR: RDBE got bad packet length")
				continue

			}

			tags := map[string]string{
				"id":      id,
				"pcalifx": fmt.Sprintf("%d", pack.PcalIfx),
				"rawifx":  fmt.Sprintf("%d", pack.RawIfx),
			}
			fields := map[string]interface{}{
				"readtime":  pack.ReadTime,
				"epochref":  pack.EpochRef,
				"epochsec":  pack.EpochSec,
				"interval":  pack.Interval,
				"mu":        pack.Mu,
				"sigma":     pack.Sigma,
				"ppsoffset": pack.PpsOffset,
				"gpsoffset": pack.GpsOffset,
			}
			u.acc.AddFields("rdbe", fields, tags, time.Now())

			if u.SaveRaw {
				//TODO: store values in a better way?
				rawfields := map[string]interface{}{
					"rawsamples": pack.RawSamples,
					"tsyson":     pack.TsysOn,
					"tsysoff":    pack.TsysOff,
					"pcalsin":    pack.PcalSin,
					"pcalcos":    pack.PcalCos,
					"statstr":    pack.StatStr,
				}
				u.acc.AddFields("rdbe_raw", rawfields, tags, time.Now())
			}
		}
	}
}

func init() {
	inputs.Add("rdbe", func() telegraf.Input {
		return &Rdbe{}
	})
}
