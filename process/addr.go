package process

import (
	"encoding/binary"
	"net"
	"reflect"
	"runtime"
	"unsafe"
)

// ParseIPAddress parses the given string in to an IPAddress
func ParseIPAddress(s string) IPAddress {
	v4, ok := fastParseIPV4(s)
	if !ok { // Slow path or for ipv6
		ip := net.ParseIP(s)

		return IPAddress{
			High: int64(binary.LittleEndian.Uint64(ip[0:8])),
			Low:  int64(binary.LittleEndian.Uint64(ip[8:])),
		}
	}

	return IPAddress{
		High: 0,
		Low:  int64(v4)<<32 | 255<<24 | 255<<16,
	}
}

// IsV4 returns true if this is a v4 ip address
func (m *IPAddress) IsV4() bool {
	return m.High == 0 && byte(m.Low) == 0 && byte(m.Low>>8) == 0 && byte(m.Low>>16) == 255 && byte(m.Low>>24) == 255
}

// String returns the string representation of this IP Address
func (m *IPAddress) String() string {
	var s string
	m.AsIP(func(ip net.IP) {
		s = ip.String()
	})
	return s
}

// V16 returns the IP address as 16 bytes
func (m *IPAddress) V16() [16]byte {
	var buf [16]byte

	binary.LittleEndian.PutUint64(buf[0:], uint64(m.High))
	binary.LittleEndian.PutUint64(buf[8:], uint64(m.Low))

	return buf
}

// V4 returns the IP address as 4 bytes
func (m *IPAddress) V4() [4]byte {
	var buf [4]byte

	binary.LittleEndian.PutUint32(buf[0:], uint32(m.Low>>32))

	return buf
}

// AsIP runs the given callback with this IP as a `net.IP`. The `net.IP` value is an unsafe version of this value
// and will only be valid for the duration of the callback function.  Parts of the `net.IP` should not be referenced
// after the function returns
//
// This is done to avoid allocating a slice for a true `net.IP`
func (m *IPAddress) AsIP(fn func(ip net.IP)) {
	v16 := m.V16()
	slice := v16[:]
	hdr := *(*reflect.SliceHeader)(unsafe.Pointer(&slice))
	ip := *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: hdr.Data,
		Len:  hdr.Len,
		Cap:  hdr.Len,
	}))

	fn(ip)

	// Ensure that the underlying slice is not garbage collected before the function call returns
	runtime.KeepAlive(slice)
}

// fastParseIPV4 is a clone of net.ParseIP but it doesn't allocate
// returns the ipv4 address as an int32 and an ok value if it was successful in parsing the value as ipv4
func fastParseIPV4(s string) (int32, bool) {
loop:
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '.':
			break loop
		case ':':
			return 0, false
		}
	}

	var p [net.IPv4len]byte
	for i := 0; i < net.IPv4len; i++ {
		if len(s) == 0 {
			// Missing octets.
			return 0, false
		}
		if i > 0 {
			if s[0] != '.' {
				return 0, false
			}
			s = s[1:]
		}
		n, c, ok := dtoi(s)
		if !ok || n > 0xFF {
			return 0, false
		}
		s = s[c:]
		p[i] = byte(n)
	}
	if len(s) != 0 {
		return 0, false
	}
	return int32(binary.LittleEndian.Uint32(p[:])), true
}

// Bigger than we need, not too big to worry about overflow
const big = 0xFFFFFF

// Decimal to integer.
// Returns number, characters consumed, success.
func dtoi(s string) (n int, i int, ok bool) {
	n = 0
	for i = 0; i < len(s) && '0' <= s[i] && s[i] <= '9'; i++ {
		n = n*10 + int(s[i]-'0')
		if n >= big {
			return big, i, false
		}
	}
	if i == 0 {
		return 0, 0, false
	}
	return n, i, true
}
