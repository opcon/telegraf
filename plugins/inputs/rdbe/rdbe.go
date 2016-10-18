package rdbe

import (
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

  ## Extra tags should be added
  ## eg.
  [inputs.cpu.tags]
    antenna = "gs"
    tag2 = "bar"
`

func (u *UdpListener) SampleConfig() string {
	return sampleConfig
}

func (u *UdpListener) Description() string {
	return "RDBE UDP Multicast listener"
}

// All the work is done in the Start() function, so this is just a dummy
// function.
func (u *UdpListener) Gather(_ telegraf.Accumulator) error {
	return nil
}

func (u *UdpListener) SetParser(parser parsers.Parser) {
	u.parser = parser
}

func (u *UdpListener) Start(acc telegraf.Accumulator) error {
	u.Lock()
	defer u.Unlock()

	u.acc = acc
	u.in = make(chan []byte, u.AllowedPendingMessages)
	u.done = make(chan struct{})
	u.listeners = make(map[string]*net.UDPConn)

	for _, id := range u.DeviceIds {
		u.wg.Add(1)
		go u.udpListen(id)
	}

	u.wg.Add(1)
	go u.udpParser()

	log.Printf("Started RDBE listener service on %s\n", u.ServiceAddress)
	return nil
}

func (u *UdpListener) Stop() {
	u.Lock()
	defer u.Unlock()
	close(u.done)
	u.wg.Wait()
	close(u.in)
	log.Println("Stopped UDP listener service on ", u.ServiceAddress)
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
	bid = byte(id[0])
	if bid < 'a' || bid > 'z' {
		log.Fatal("bad rdbe id %s", id)
	}
	addr := fmt.Sprintf("239.0.2.%d", (bid-'a'+1)*10)
	port := 20021 + int(bid-'a')
	return fmt.Sprintf("%s:%d", addr, port)
}

func (u *UdpListener) rdbeListen(id string) {
	defer u.wg.Done()

	addr, _ := net.ResolveUDPAddr("udp", broadcastAddress(id))
	conn, err := net.ListenMulticastUDP("udp", nil, addr)
	defer conn.Close()
	if err != nil {
		log.Fatalf("ERROR: ListenUDP - %s", err)
	}
	log.Printf("RDBE server listening to %s", listener.LocalAddr().String())

	var pack rdbepacket
	buf := make([]byte, UDP_MAX_PACKET_SIZE)

	for {
		select {
		case <-u.done:
			return
		default:
			n, err := conn.Read(buff)
			if err != nil {
				log.Printf("ERROR: %s\n", err.Error())
				continue
			}

			reader := bytes.NewReader(buff)
			err = binary.Read(reader, binary.BigEndian, &pack)
			if err != nil {
				log.Printf("ERROR: %s\n", err)
				continue
			}

			if n != int(pack.PacketSize) {
				log.Println("ERROR: RDBE got bad packet length")
				continue
			}

			u.acc.AddFields(m.Name(), m.Fields(), m.Tags(), m.Time())
		}
	}
}

func init() {
	inputs.Add("udp_listener", func() telegraf.Input {
		return &UdpListener{}
	})
}
