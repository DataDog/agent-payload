package process

import (
	"encoding/binary"
	"net"
)

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
