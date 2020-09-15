package process

import (
	"encoding/binary"
	"net"
)

// ParseIPAddress parses the given string in to an IPAddress
func ParseIPAddress(s string) IPAddress {
	ip := net.ParseIP(s)

	return IPAddress{
		High: int64(binary.LittleEndian.Uint64(ip[0:8])),
		Low:  int64(binary.LittleEndian.Uint64(ip[8:])),
	}
}

// ToIP returns the `net.IP` version of this value
func (m *IPAddress) ToIP() net.IP {
	var buf [16]byte

	binary.LittleEndian.PutUint64(buf[0:], uint64(m.High))
	binary.LittleEndian.PutUint64(buf[8:], uint64(m.Low))
	return buf[:]
}

// String returns the string representation of this IP Address
func (m *IPAddress) String() string {
	return m.ToIP().String()
}
