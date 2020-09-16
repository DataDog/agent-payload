package process

import (
	"encoding/binary"
	"net"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseIPAddress(t *testing.T) {
	val := "10.12.255.123"
	ip := ParseIPAddress(val)
	var buf [16]byte

	binary.LittleEndian.PutUint64(buf[0:], uint64(ip.High))
	binary.LittleEndian.PutUint64(buf[8:], uint64(ip.Low))

	assert.Equal(t, val, ip.String())
}

func TestIPAddress_String(t *testing.T) {
	addr := ParseIPAddress("10.140.1.243")

	assert.Equal(t, "10.140.1.243", addr.String())
}

func TestIPAddress_V4(t *testing.T) {
	addr := ParseIPAddress("10.140.1.243")

	assert.Equal(t, [4]byte{10, 140, 1, 243}, addr.V4())
}

func TestIPAddress_IsV4(t *testing.T) {
	for _, s := range []string{
		"127.0.0.1",
		"10.140.1.243",
		"1.1.1.1",
		"255.255.255.255",
		"0.0.0.0",
	} {
		addr := ParseIPAddress(s)
		assert.True(t, addr.IsV4(), s)
	}

	for _, s := range []string{
		"2001:4860:0:2001::68",
	} {
		addr := ParseIPAddress(s)
		assert.False(t, addr.IsV4(), s)
	}
}

func TestIPAddress_AsIP(t *testing.T) {
	for _, ip := range []string{
		"127.0.0.1",
		"10.140.1.243",
		"1.1.1.1",
		"255.255.255.255",
	} {
		addr := ParseIPAddress(ip)

		var s string
		addr.AsIP(func(ip net.IP) {
			s = ip.String()
		})

		assert.Equal(t, ip, s)
	}
}

func BenchmarkIPAddress_AsIP(b *testing.B) {
	addr := ParseIPAddress("127.0.0.1")

	b.ReportAllocs()
	b.ResetTimer()

	var v bool

	for i := 0; i < b.N; i++ {
		addr.AsIP(func(ip net.IP) {
			v = ip.IsLoopback()
		})
	}

	runtime.KeepAlive(v)
}

func BenchmarkParseIPAddress(b *testing.B) {
	b.Run("ipv4", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		var addr IPAddress

		for i := 0; i < b.N; i++ {
			addr = ParseIPAddress("10.1.1.1")
		}

		runtime.KeepAlive(addr)
	})
}
