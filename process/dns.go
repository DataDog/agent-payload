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
func IterateDNS(buf []byte, ip string, cb func(i, total int, entry string) bool) error {
	if len(buf) == 0 || ip == "" {
		return nil
	}

	switch buf[0] {
	case dnsVersion1:
		return iterateDNSV1(buf, ip, cb)
	}
	return nil
}

// UnsafeIterateDNS invokes the callback function for each DNS entry for the given IP in the given buffer.
// Each entry is a the slice from the overall buffer.  It should be copied before use
func UnsafeIterateDNS(buf []byte, ip string, cb func(i, total int, entry []byte) bool) error {
	if len(buf) == 0 || ip == "" {
		return nil
	}

	switch buf[0] {
	case dnsVersion1:
		return unsafeIterateDNSV1(buf, ip, cb)
	}

	return nil
}
