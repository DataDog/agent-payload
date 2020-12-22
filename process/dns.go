package process

type DNSEncoder interface {
	Encode(dns map[string]*DNSEntry) []byte
}

const dnsVersion1 byte = 1

// GetDNS gets the DNS entries for the given IP from the given buffer
func GetDNS(buf []byte, ip string) (string, []string) {
	if len(buf) == 0 || ip == "" {
		return "", nil
	}

	switch buf[0] {
	case dnsVersion1:
		return getV1(buf, ip)
	}

	return "", nil
}

func getDNSNames(buf []byte) []string {
	if len(buf) == 0 {
		return nil
	}

	switch buf[0] {
	case dnsVersion1:
		return getDNSNamesV1(buf)
	}
	return nil
}

// IterateDNS invokes the callback function for each DNS entry for the given IP in the given buffer
func IterateDNS(buf []byte, ip string, cb func(i, total int, entry string) bool) {
	if len(buf) == 0 || ip == "" {
		return
	}

	switch buf[0] {
	case dnsVersion1:
		iterateDNSV1(buf, ip, cb)
	}
}
