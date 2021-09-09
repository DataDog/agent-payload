package process

import (
	"io/ioutil"
	"path"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestV2DomainDatabaseEncoding(t *testing.T) {
	dnsdb := []string{
		"foo.com",
		"service.example.com",
		"service2.example.com",
		"app.example.com",
		"bar.com",
	}
	encoder := NewV2DNSEncoder()
	buf, offsets, err := encoder.EncodeDomainDatabase(dnsdb)
	assert.Nil(t, err)

	decoded := getDNSNameListV2(buf)
	for idx, s := range dnsdb {
		assert.Equal(t, s, decoded[idx])

		byIndex, err := getDNSNameFromListByIndex(buf, idx)
		assert.Nil(t, err)
		assert.Equal(t, s, byIndex)

		byOffset, err := getDNSNameFromListByOffset(buf, int(offsets[idx]))
		assert.Nil(t, err)
		assert.Equal(t, s, byOffset)
	}

	// test out of bounds
	_, err = getDNSNameFromListByIndex(buf, 7)
	assert.Error(t, err)

	// test off of the end
	_, err = getDNSNameFromListByOffset(buf, len(buf)+2)
	assert.Error(t, err)

}
func indexOf(val string, db []string) int32 {
	for p, v := range db {
		if v == val {
			return int32(p)
		}
	}
	return -1
}

func TestV2EncodeDNS(t *testing.T) {
	dns := make(map[string]*DNSDatabaseEntry)

	dnsdb := []string{
		"foo.com",
		"service.example.com",
		"service2.example.com",
		"app.example.com",
		"bar.com",
	}

	dns["10.128.98.75"] = &DNSDatabaseEntry{NameOffsets: []int32{indexOf("service.example.com", dnsdb), indexOf("service2.example.com", dnsdb)}}
	dns["10.128.99.240"] = &DNSDatabaseEntry{NameOffsets: []int32{indexOf("service.example.com", dnsdb)}}
	dns["34.231.44.115"] = &DNSDatabaseEntry{NameOffsets: []int32{indexOf("app.example.com", dnsdb)}}

	encoder := NewV2DNSEncoder()
	encodedDatabase, offsets, err := encoder.EncodeDomainDatabase(dnsdb)
	buf, err := encoder.EncodeMapped(dns, offsets)
	assert.Nil(t, err)

	decodedDatabase := getDNSNameListV2(encodedDatabase)

	assert.Equal(t, len(dnsdb), len(decodedDatabase))

	assertDNSV2Equal(t, []string{"service.example.com", "service2.example.com"}, buf, encodedDatabase, "10.128.98.75")
	assertDNSV2Equal(t, []string{"service.example.com"}, buf, encodedDatabase, "10.128.99.240")
	assertDNSV2Equal(t, []string{"app.example.com"}, buf, encodedDatabase, "34.231.44.115")
	assertDNSV2Equal(t, nil, buf, encodedDatabase, "134.231.44.115")
	assertDNSV2Equal(t, nil, buf, encodedDatabase, "1.1.1.1")

}

func TestV2EncodeDNS_Empty(t *testing.T) {
	dns := make(map[string]*DNSDatabaseEntry)

	encoder := NewV2DNSEncoder()
	buf, err := encoder.EncodeMapped(dns, nil)

	assert.Nil(t, err)
	assert.Empty(t, buf)
	assertDNSV2Equal(t, nil, buf, nil, "1.1.1.1")

	emptydb := make([]byte, 0)
	names, err := getDNSNames(emptydb)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(names))
}

func TestV2EncodeDNS_NoNames(t *testing.T) {
	dns := make(map[string]*DNSDatabaseEntry)

	dns["10.128.98.75"] = &DNSDatabaseEntry{NameOffsets: []int32{}}
	dns["10.128.99.240"] = &DNSDatabaseEntry{}

	encoder := NewV2DNSEncoder()
	buf, err := encoder.EncodeMapped(dns, nil)

	assert.Nil(t, err)

	assert.Empty(t, buf)
	assertDNSV2Equal(t, nil, buf, nil, "10.128.98.75")
	assertDNSV2Equal(t, nil, buf, nil, "10.128.99.240")

}

func TestV2EncodeDNS_SampleData(t *testing.T) {
	sampleFiles := []string{
		"testdata/dns/samples.txt",
		"testdata/dns/big_ips.txt",
		"testdata/dns/big_entries.txt",
	}

	for _, sampleFile := range sampleFiles {
		t.Run(path.Base(sampleFile), func(t *testing.T) {
			samples, stringdb := readTestDnsV2(t, sampleFile)

			encoder := NewV2DNSEncoder()

			encodedDb, indexToOffset, err := encoder.EncodeDomainDatabase(stringdb)

			assert.Nil(t, err)

			for _, sample := range samples {
				buf, _ := encoder.EncodeMapped(sample, indexToOffset)

				for ip, entry := range sample {
					// the entry we read from file is stored by index.  Get the names
					// by index, and use that to compare
					var expected []string
					for _, idx := range entry.NameOffsets {
						expected = append(expected, stringdb[idx])
					}

					assertDNSV2Equal(t, expected, buf, encodedDb, ip)
				}
			}
		})

	}
}

func TestV2DncodeDNS_SampleData(t *testing.T) {
	sampleFiles := []string{
		"testdata/dns/samples.txt",
		"testdata/dns/big_ips.txt",
		"testdata/dns/big_entries.txt",
	}

	for _, sampleFile := range sampleFiles {
		t.Run(path.Base(sampleFile), func(t *testing.T) {
			_, sampledb := readTestDnsV2(t, sampleFile)

			encoder := NewV2DNSEncoder()
			buf, indexToOffset, err := encoder.EncodeDomainDatabase(sampledb)
			assert.Nil(t, err)

			decodedDb := getDNSNameListV2(buf)
			assert.Equal(t, sampledb, decodedDb)

			for idx, name := range sampledb {
				decoded, err := getDNSNameFromListByIndex(buf, idx)
				assert.Nil(t, err)
				assert.Equal(t, name, decoded)
				decoded, err = getDNSNameFromListByOffset(buf, int(indexToOffset[idx]))
				assert.Nil(t, err)
				assert.Equal(t, name, decoded)
			}
		})

	}
}

func BenchmarkDNSV2Decode(b *testing.B) {
	sampleFiles := []string{
		"testdata/dns/samples.txt",
		"testdata/dns/big_ips.txt",
		"testdata/dns/big_entries.txt",
	}

	encoder := NewV2DNSEncoder()

	for _, sampleFile := range sampleFiles {
		samples, dnsdb := readTestDnsV2(b, sampleFile)
		_, indexToOffset, _ := encoder.EncodeDomainDatabase(dnsdb)

		b.Run(path.Base(sampleFile), func(b *testing.B) {
			bufs := make([][]byte, len(samples))

			for i, dns := range samples {
				bufs[i], _ = encoder.EncodeMapped(dns, indexToOffset)
			}

			b.ReportAllocs()
			b.ResetTimer()

			var s []int32

			for i := 0; i < b.N; i++ {
				for i, dns := range samples {
					for ip := range dns {
						_, s, _ = GetDNSV2(bufs[i], ip)
					}
				}
			}

			runtime.KeepAlive(s)
		})
	}
}

func BenchmarkDNSV2Encode(b *testing.B) {
	sampleFiles := []string{
		"testdata/dns/samples.txt",
		"testdata/dns/big_ips.txt",
		"testdata/dns/big_entries.txt",
	}

	encoder := NewV2DNSEncoder()

	for _, sampleFile := range sampleFiles {
		samples, dnsdb := readTestDnsV2(b, sampleFile)
		_, indexToOffset, _ := encoder.EncodeDomainDatabase(dnsdb)

		b.Run(path.Base(sampleFile), func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()

			var buf []byte
			var count int64

			for i := 0; i < b.N; i++ {
				for _, dns := range samples {
					buf, _ = encoder.EncodeMapped(dns, indexToOffset)
					count += int64(len(buf))
				}
			}

			b.ReportMetric(float64(count)/float64(b.N), "bytes")
			runtime.KeepAlive(buf)
		})
	}
}

func appendToDatabase(name string, present *map[string]int32, db *[]string) int32 {
	if idx, ok := (*present)[name]; ok {
		return idx
	}
	len := int32(len(*db))
	*db = append(*db, name)
	(*present)[name] = len
	return len
}

func readTestDnsV2(t require.TestingT, filename string) ([]map[string]*DNSDatabaseEntry, []string) {
	buf, err := ioutil.ReadFile(filename)
	require.NoError(t, err)

	var maps []map[string]*DNSDatabaseEntry
	namedb := make([]string, 0)
	namemap := make(map[string]int32)

	for _, line := range strings.Split(string(buf), "\n") {
		entries := strings.Split(line, "|")
		data := make(map[string]*DNSDatabaseEntry)

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

			filtered := make([]int32, 0)
			for _, name := range names {
				if len(name) > 0 {
					idx := appendToDatabase(name, &namemap, &namedb)
					filtered = append(filtered, idx)
				}
			}

			data[ip] = &DNSDatabaseEntry{NameOffsets: filtered}
		}

		maps = append(maps, data)
	}

	return maps, namedb
}

func assertDNSV2Equal(t *testing.T, expected []string, buf []byte, dnsdb []byte, key string) {
	name, names, err := GetDNSV2(buf, key)

	assert.Nil(t, err)
	switch len(expected) {
	case 0:
		assert.Equal(t, int32(-1), name)
		assert.Empty(t, names)
	default:
		namestr, err := getDNSNameFromListByOffset(dnsdb, int(name))
		assert.Nil(t, err)
		assert.Equal(t, expected[0], namestr)

		for arrayindex, offset := range names {
			namestr, err := getDNSNameFromListByOffset(dnsdb, int(offset))
			assert.Nil(t, err)
			assert.Equal(t, expected[arrayindex+1], namestr)

		}
	}

	var iterValues []int32
	IterateDNSV2(buf, key, func(i, total int, entry int32) bool {
		iterValues = append(iterValues, entry)
		return true
	})

	var unsafeIterValues []int32
	UnsafeIterateDNSV2(buf, key, func(i, total int, entry int32) bool {
		unsafeIterValues = append(unsafeIterValues, entry)
		return true
	})

	var truncatedValues []int32
	IterateDNSV2(buf, key, func(i, total int, entry int32) bool {
		if i == total-1 {
			return false
		}
		truncatedValues = append(truncatedValues, entry)
		return true
	})

	switch len(iterValues) {
	case 0:
		assert.Equal(t, int32(-1), name)
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
