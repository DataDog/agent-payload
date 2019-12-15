package process

type DNSEncoder interface {
	Encode(dns map[string]*DNSEntry) []byte
}

const dnsVersion1 byte = 1

func getDNS(buf []byte, ip string) (string, []string) {
	if len(buf) == 0 || ip == "" {
		return "", nil
	}

	switch buf[0] {
	case dnsVersion1:
		return getV1(buf, ip)
	}

	return "", nil
}
