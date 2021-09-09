package process

type DNSEncoder interface {
	Encode(dns map[string]*DNSEntry) ([]byte, error)
	EncodeMapped(dns map[string]*DNSDatabaseEntry, indexToOffset []int32) ([]byte, error)
	EncodeDomainDatabase(names []string) ([]byte, []int32, error)
}

const dnsVersion1 byte = 1
const dnsVersion2 byte = 2
