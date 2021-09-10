package process

func (m *CollectorConnections) GetHostTags(host *Host) []string {
	return m.GetTags(int(host.TagIndex))
}

func (m *CollectorConnections) IterateHostTags(host *Host, cb func(i, total int, tag string) bool) {
	iterateTags(m.EncodedTags, int(host.TagIndex), cb)
}

func (m *CollectorConnections) GetResourceTags(resource *ResourceMetadata) []string {
	return m.GetTags(int(resource.TagIndex))
}

func (m *CollectorConnections) IterateResourceTags(resource *ResourceMetadata, cb func(i, total int, tag string) bool) {
	iterateTags(m.EncodedTags, int(resource.TagIndex), cb)
}

func (m *CollectorConnections) GetTags(tagIndex int) []string {
	return getTags(m.EncodedTags, tagIndex)
}

// GetDNS returns the DNS entries for the given addr.
// The first argument returned is the first DNS entry followed by any additional name resolutions.  Most IPs will
// have a single resolution so this dual format allows us to avoid allocations for the common case.  If there are
// multiple name resolutions, there is no implied priority between the dual values
func (m *CollectorConnections) GetDNS(addr *Addr) (string, []string, error) {
	return GetDNS(m.EncodedDNS, addr.Ip)
}

// IterateDNS iterates over all of the DNS entries for the given addr, invoking the provided callback for each one
func (m *CollectorConnections) IterateDNS(addr *Addr, cb func(i, total int, entry string) bool) {
	IterateDNS(m.EncodedDNS, addr.Ip, cb)
}

// GetDNSNames returns all the DNS entries
func (m *CollectorConnections) GetDNSNames() ([]string, error) {
	return getDNSNames(m.EncodedDNS)
}

// GetDNSV2 returns the DNS entries for the given addr.
// The first argument returned is the first DNS entry followed by any additional name resolutions.  Most IPs will
// have a single resolution so this dual format allows us to avoid allocations for the common case.  If there are
// multiple name resolutions, there is no implied priority between the dual values
// the return values are indexes into the dns database (GetDNSDatabase)

func (m *CollectorConnections) GetDNSV2(addr *Addr) (int32, []int32, error) {
	return GetDNSV2(m.EncodedDnsLookups, addr.Ip)
}

// IterateDNSV2 iterates over all of the DNS entries for the given addr, invoking the provided callback for each one
func (m *CollectorConnections) IterateDNSV2(addr *Addr, cb func(i, total int, entry int32) bool) {
	IterateDNSV2(m.EncodedDnsLookups, addr.Ip, cb)
}

// GetDNSNamesV2 returns all the DNS entries
func (m *CollectorConnections) GetDNSNamesV2() []string {
	return getDNSNameListV2(m.EncodedDomainDatabase)
}
