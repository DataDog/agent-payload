package process

import (
	"io/ioutil"
	"path"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

func TestV1EncodeDNS(t *testing.T) {
	dns := make(map[string]*DNSEntry)

	dns["10.128.98.75"] = &DNSEntry{Names: []string{"service.example.com", "service2.example.com"}}
	dns["10.128.99.240"] = &DNSEntry{Names: []string{"service.example.com"}}
	dns["34.231.44.115"] = &DNSEntry{Names: []string{"app.example.com"}}

	encoder := NewV1DNSEncoder()
	buf, err := encoder.Encode(dns)

	assert.Nil(t, err)

	assertDNSEqual(t, []string{"service.example.com", "service2.example.com"}, buf, "10.128.98.75")
	assertDNSEqual(t, []string{"service.example.com"}, buf, "10.128.99.240")
	assertDNSEqual(t, []string{"app.example.com"}, buf, "34.231.44.115")
	assertDNSEqual(t, nil, buf, "134.231.44.115")
	assertDNSEqual(t, nil, buf, "1.1.1.1")

	names, err := getDNSNames(buf)
	assert.Nil(t, err)
	assert.Equal(t, 3, len(names))
}

func TestV1EncodeDNS_Empty(t *testing.T) {
	dns := make(map[string]*DNSEntry)

	encoder := NewV1DNSEncoder()
	buf, err := encoder.Encode(dns)

	assert.Nil(t, err)
	assert.Empty(t, buf)
	assertDNSEqual(t, nil, buf, "1.1.1.1")

	names, err := getDNSNames(buf)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(names))
}

func TestV1EncodeDNS_NoNames(t *testing.T) {
	dns := make(map[string]*DNSEntry)

	dns["10.128.98.75"] = &DNSEntry{Names: []string{}}
	dns["10.128.99.240"] = &DNSEntry{}

	encoder := NewV1DNSEncoder()
	buf, err := encoder.Encode(dns)

	assert.Nil(t, err)

	assert.Empty(t, buf)
	assertDNSEqual(t, nil, buf, "10.128.98.75")
	assertDNSEqual(t, nil, buf, "10.128.99.240")

	names, err := getDNSNames(buf)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(names))
}

func TestV1EncodeDNS_SampleData(t *testing.T) {
	sampleFiles := []string{
		"testdata/dns/samples.txt",
		"testdata/dns/big_ips.txt",
		"testdata/dns/big_entries.txt",
	}

	for _, sampleFile := range sampleFiles {
		t.Run(path.Base(sampleFile), func(t *testing.T) {
			samples := readTestDns(t, sampleFile)

			encoder := NewV1DNSEncoder()

			for _, sample := range samples {
				buf, _ := encoder.Encode(sample)

				for ip, entry := range sample {
					assertDNSEqual(t, entry.Names, buf, ip)
				}
			}
		})
	}
}

func BenchmarkDNSDecode(b *testing.B) {
	sampleFiles := []string{
		"testdata/dns/samples.txt",
		"testdata/dns/big_ips.txt",
		"testdata/dns/big_entries.txt",
	}

	encoder := NewV1DNSEncoder()

	for _, sampleFile := range sampleFiles {
		samples := readTestDns(b, sampleFile)

		b.Run(path.Base(sampleFile), func(b *testing.B) {
			bufs := make([][]byte, len(samples))

			for i, dns := range samples {
				bufs[i], _ = encoder.Encode(dns)
			}

			b.ReportAllocs()
			b.ResetTimer()

			var s []string

			for i := 0; i < b.N; i++ {
				for i, dns := range samples {
					for ip := range dns {
						_, s, _ = GetDNS(bufs[i], ip)
					}
				}
			}

			runtime.KeepAlive(s)
		})
	}
}

func BenchmarkDNSEncode(b *testing.B) {
	sampleFiles := []string{
		"testdata/dns/samples.txt",
		"testdata/dns/big_ips.txt",
		"testdata/dns/big_entries.txt",
	}

	encoder := NewV1DNSEncoder()

	for _, sampleFile := range sampleFiles {
		samples := readTestDns(b, sampleFile)

		b.Run(path.Base(sampleFile), func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()

			var buf []byte
			var count int64

			for i := 0; i < b.N; i++ {
				for _, dns := range samples {
					buf, _ = encoder.Encode(dns)
					count += int64(len(buf))
				}
			}

			b.ReportMetric(float64(count)/float64(b.N), "bytes")
			runtime.KeepAlive(buf)
		})
	}
}

func readTestDns(t require.TestingT, filename string) []map[string]*DNSEntry {
	buf, err := ioutil.ReadFile(filename)
	require.NoError(t, err)

	var maps []map[string]*DNSEntry
	for _, line := range strings.Split(string(buf), "\n") {
		entries := strings.Split(line, "|")
		data := make(map[string]*DNSEntry)

		for _, entry := range entries {
			if len(entry) == 0 {
				continue
			}

			idx := strings.IndexByte(entry, ':')
			if idx == -1 {
				continue
			}

			ip := entry[:idx]
			names := strings.Split(entry[idx+1:], ",")

			filtered := names[:0]
			for _, name := range names {
				if len(name) > 0 {
					filtered = append(filtered, name)
				}
			}

			data[ip] = &DNSEntry{Names: filtered}
		}

		maps = append(maps, data)
	}

	return maps
}

func assertDNSEqual(t *testing.T, expected []string, buf []byte, key string) {
	name, names, err := GetDNS(buf, key)

	assert.Nil(t, err)
	switch len(expected) {
	case 0:
		assert.Empty(t, name)
		assert.Empty(t, names)
	case 1:
		assert.Equal(t, expected[0], name)
		assert.Empty(t, names)
	default:
		assert.Equal(t, expected[0], name)
		assert.Equal(t, expected[1:], names)
	}

	var iterValues []string
	IterateDNS(buf, key, func(i, total int, entry string) bool {
		iterValues = append(iterValues, entry)
		return true
	})

	var unsafeIterValues []string
	UnsafeIterateDNS(buf, key, func(i, total int, entry []byte) bool {
		unsafeIterValues = append(unsafeIterValues, string(entry))
		return true
	})

	var truncatedValues []string
	IterateDNS(buf, key, func(i, total int, entry string) bool {
		if i == total-1 {
			return false
		}
		truncatedValues = append(truncatedValues, entry)
		return true
	})

	switch len(iterValues) {
	case 0:
		assert.Empty(t, name)
		assert.Empty(t, names)

		assert.Empty(t, truncatedValues)
	case 1:
		assert.Equal(t, name, iterValues[0])
		assert.Equal(t, name, unsafeIterValues[0])
		assert.Empty(t, truncatedValues)
	default:
		assert.Equal(t, name, iterValues[0])
		assert.Equal(t, names, iterValues[1:])

		assert.Equal(t, name, unsafeIterValues[0])
		assert.Equal(t, names, unsafeIterValues[1:])

		assert.Equal(t, iterValues[0:len(iterValues)-1], truncatedValues)
	}

}
