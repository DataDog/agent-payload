package process

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"path"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

var (
	garbageData = "0109000cb4020400012223247687018801990100020c332e3233332e3134342e353201000d332e3233332e3134382e31323001380000050d35322e3231362e37362e31303801160d332e3233332e3134382e31313801000d33342e3132302e31352e31373301750b35322e3231372e37392e340200000d332e3233332e3134372e3139370100010d332e3233332e3134372e323039010000010c3137322e32362e34382e313001d301030d332e3233332e3134372e31393801000d33342e3130372e3137322e323301570e33342e3131372e3231382e3232370199010000001570726f636573732e64617461646f6768712e636f6d21656c61737469636d61707265647563652e73332e616d617a6f6e6177732e636f6d1e362d33312d312d6170702e6167656e742e64617461646f6768712e636f6d1d362d33312d312d6170702e6167656e742e64617461646f6768712e6575236167656e742d687474702d696e74616b652e6c6f67732e64617461646f6768712e65751470726f636573732e64617461646f6768712e6575246167656e742d687474702d696e74616b652e6c6f67732e64617461646f6768712e636f6d1c69702d3137322d32362d34382d31302e656d722e70726f642e646f67216177733135372d6c6f67732d70726f642e73332e616d617a6f6e6177732e636f6d2164642d656d722d6c6f67732d70726f642e73332e616d617a6f6e6177732e636f6d"
	//myData    = "0109000cb4020400012324257889018a019a0100020d332e3233332e3134382e3132300188010c332e3233332e3134342e353201a7010000050d332e3233332e3134372e31393701000d35322e3231362e37362e31303801160d332e3233332e3134382e31313801000b35322e3231372e37392e3402cc01ee010d33342e3132302e31352e313733019002010d332e3233332e3134372e323039010000010c3137322e32362e34382e3130016b030d332e3233332e3134372e31393801000e33342e3131372e3231382e32323701380d33342e3130372e3137322e3233014d1570726f636573732e64617461646f6768712e636f6d21656c61737469636d61707265647563652e73332e616d617a6f6e6177732e636f6d1470726f636573732e64617461646f6768712e65751d362d33312d312e6170702e6167656e742e64617461646f6768712e65751c69702d3137322d32362d34382d31302e656d722e70726f642e646f671e362d33312d312d6170702e6167656e742e64617461646f6768712e636f6d246167656e742d687474702d696e74616b652e6c6f67732e64617461646f6768712e636f6d2164642d656d722d6c6f67732d70726f642e73332e616d617a6f6e6177732e636f6d216177733135372d6c6f67732d70726f642e73332e616d617a6f6e6177732e636f6d236167656e742d687474702d696e74616b652e6c6f67732e64617461646f6768712e6575"
	//mydata2   = "0109000cb4020400012324257789018a019b0100020d332e3233332e3134382e31323001d2010c332e3233332e3134342e353201f1010000050b35322e3231372e37392e340200220d33342e3132302e31352e31373301440d35322e3231362e37362e31303801680d332e3233332e3134372e313937018a010d332e3233332e3134382e313138018a01010d332e3233332e3134372e323039018a0100010c3137322e32362e34382e313001a001030e33342e3131372e3231382e32323701bd010d33342e3130372e3137322e32330196020d332e3233332e3134372e313938018a012164642d656d722d6c6f67732d70726f642e73332e616d617a6f6e6177732e636f6d216177733135372d6c6f67732d70726f642e73332e616d617a6f6e6177732e636f6d236167656e742d687474702d696e74616b652e6c6f67732e64617461646f6768712e657521656c61737469636d61707265647563652e73332e616d617a6f6e6177732e636f6d1570726f636573732e64617461646f6768712e636f6d1c69702d3137322d32362d34382d31302e656d722e70726f642e646f671470726f636573732e64617461646f6768712e65751e362d33312d312d6170702e6167656e742e64617461646f6768712e636f6d246167656e742d687474702d696e74616b652e6c6f67732e64617461646f6768712e636f6d1d362d33312d312e6170702e6167656e742e64617461646f6768712e6575
)

func TestMatchGarbageData(t *testing.T) {
	baddata, err := hex.DecodeString(garbageData)
	assert.Nil(t, err)

	Decode(baddata)

	dns := make(map[string]*DNSEntry)

	dns["3.233.148.120"] = &DNSEntry{Names: []string{"6-31-1-app.agent.datadoghq.com"}}
	dns["3.233.144.52"] = &DNSEntry{Names: []string{"agent-http-intake.logs.datadoghq.com"}}

	dns["3.233.147.197"] = &DNSEntry{Names: []string{"process.datadoghq.com"}}
	dns["52.217.79.4"] = &DNSEntry{Names: []string{"aws157-logs-prod.s3.amazonaws.com", "dd-emr-logs-prod.s3.amazonaws.com"}}
	dns["34.120.15.173"] = &DNSEntry{Names: []string{"agent-http-intake.logs.datadoghq.eu"}}
	dns["3.233.148.118"] = &DNSEntry{Names: []string{"process.datadoghq.com"}}
	dns["52.216.76.108"] = &DNSEntry{Names: []string{"elasticmapreduce.s3.amazonaws.com"}}

	dns["3.233.147.209"] = &DNSEntry{Names: []string{"process.datadoghq.com"}}

	dns["172.26.48.10"] = &DNSEntry{Names: []string{"ip-172-26-48-10.emr.prod.dog"}}

	dns["34.117.218.227"] = &DNSEntry{Names: []string{"process.datadoghq.eu"}}
	dns["34.107.172.23"] = &DNSEntry{Names: []string{"6-31-1.app.agent.datadoghq.eu"}}
	dns["3.233.147.198"] = &DNSEntry{Names: []string{"process.datadoghq.com"}}

	encoder := NewV1DNSEncoder()
	runs := 10
	results := make([][]byte, runs)
	for i := 0; i < runs; i++ {
		testbuf := encoder.Encode(dns)
		results[i] = testbuf
		if bytes.Compare(testbuf, baddata) == 0 {
			fmt.Printf("Run %d matched bad data\n", i)
		}
	}
	// now see if any of them match
	for out := 0; out < runs; out++ {
		for in := out + 1; in < runs; in++ {
			if bytes.Compare(results[out], results[in]) == 0 {
				fmt.Printf("run %d matched run %d\n", out, in)
			}
		}
	}

}
func TestGarbageData(t *testing.T) {

	baddata, err := hex.DecodeString(garbageData)
	assert.Nil(t, err)

	dns, err := Decode(baddata)
	assert.Nil(t, err)
	fmt.Printf("%v\n", dns)

	encoder := NewV1DNSEncoder()
	buf := encoder.Encode(dns)
	assert.Equal(t, baddata, buf)

	/*
		for k, v := range dns {
			fmt.Printf("IP: %s\n", k)
			for _, name := range v.Names {
				fmt.Printf("   %s\n", name)
			}
		}
	*/
	//	for k, _ := range dns {
	UnsafeIterateDNS(buf, "3.233.147.209", func(i, total int, entry []byte) bool {
		fmt.Printf("Key %s %d of %d %s\n", "3.233.147.209", i, total, string(entry))
		return true
	})
	//	}
	fmt.Printf("done\n")

}
func TestV1EncodeDecode(t *testing.T) {
	dns := make(map[string]*DNSEntry)

	dns["10.128.98.75"] = &DNSEntry{Names: []string{"service.example.com", "service2.example.com"}}
	dns["10.128.99.240"] = &DNSEntry{Names: []string{"service.example.com"}}
	dns["34.231.44.115"] = &DNSEntry{Names: []string{"app.example.com"}}

	encoder := NewV1DNSEncoder()
	buf := encoder.Encode(dns)

	decoded, err := Decode(buf)
	assert.Nil(t, err)
	fmt.Printf("%v\n", decoded)

}
func TestV1EncodeDNS(t *testing.T) {
	dns := make(map[string]*DNSEntry)

	dns["10.128.98.75"] = &DNSEntry{Names: []string{"service.example.com", "service2.example.com"}}
	dns["10.128.99.240"] = &DNSEntry{Names: []string{"service.example.com"}}
	dns["34.231.44.115"] = &DNSEntry{Names: []string{"app.example.com"}}

	encoder := NewV1DNSEncoder()
	buf := encoder.Encode(dns)

	assertDNSEqual(t, []string{"service.example.com", "service2.example.com"}, buf, "10.128.98.75")
	assertDNSEqual(t, []string{"service.example.com"}, buf, "10.128.99.240")
	assertDNSEqual(t, []string{"app.example.com"}, buf, "34.231.44.115")
	assertDNSEqual(t, nil, buf, "134.231.44.115")
	assertDNSEqual(t, nil, buf, "1.1.1.1")
	assert.Equal(t, 3, len(getDNSNames(buf)))
}

func TestV1EncodeDNS_Empty(t *testing.T) {
	dns := make(map[string]*DNSEntry)

	encoder := NewV1DNSEncoder()
	buf := encoder.Encode(dns)

	assert.Empty(t, buf)
	assertDNSEqual(t, nil, buf, "1.1.1.1")
	assert.Equal(t, 0, len(getDNSNames(buf)))
}

func TestV1EncodeDNS_NoNames(t *testing.T) {
	dns := make(map[string]*DNSEntry)

	dns["10.128.98.75"] = &DNSEntry{Names: []string{}}
	dns["10.128.99.240"] = &DNSEntry{}

	encoder := NewV1DNSEncoder()
	buf := encoder.Encode(dns)

	assert.Empty(t, buf)
	assertDNSEqual(t, nil, buf, "10.128.98.75")
	assertDNSEqual(t, nil, buf, "10.128.99.240")
	assert.Equal(t, 0, len(getDNSNames(buf)))
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
				buf := encoder.Encode(sample)

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
				bufs[i] = encoder.Encode(dns)
			}

			b.ReportAllocs()
			b.ResetTimer()

			var s []string

			for i := 0; i < b.N; i++ {
				for i, dns := range samples {
					for ip := range dns {
						_, s = GetDNS(bufs[i], ip)
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
					buf = encoder.Encode(dns)
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
	name, names := GetDNS(buf, key)

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
