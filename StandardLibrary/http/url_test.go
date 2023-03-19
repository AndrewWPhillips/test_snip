package __

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"net/url"
	"testing"
)

func TestURLParse(t *testing.T) {
	qqq, err := url.Parse("dummy://guest:pw@a.b.com:80")
	if err != nil {
		print(err)
	}
	fmt.Printf("%q %q %q\n", qqq.Hostname(), qqq.Port(), qqq.User.Username())
	fmt.Println(qqq.User.Password())
}

// Test classless internet domain routing
func TestCIDRParse(t *testing.T) {
	// 10.0.0.1/32 => IP: A000 0001  Mask: FFFF FFFF
	// 10.0.0.1/24 => IP: A000 0000  Mask: FFFF FF00
	ip, ipnet, err := net.ParseCIDR("10.0.0.1/32")
	fmt.Println(ip, ipnet, err)
	a := binary.BigEndian.Uint32(ipnet.IP)
	b := binary.BigEndian.Uint32(ipnet.Mask)
	log.Printf("%x %x\n", a, b)
}

func TestIPParse(t *testing.T) {
	ip4 := net.ParseIP("10.9.8.0").To4()
	log.Printf("%v\n", ip4)
	log.Printf("%x\n", binary.BigEndian.Uint32(ip4)+2)
	log.Printf("%x\n", binary.LittleEndian.Uint32(ip4)+2)

	ip4[3] = 2
	log.Printf("%v\n", ip4)
}
