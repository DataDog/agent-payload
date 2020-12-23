package process

func (m *CollectorConnections) GetHostTags(host *Host) []string {
	return m.GetTags(int(host.TagIndex))
}

func (m *CollectorConnections) IterateHostTags(host *Host, cb func(i, total int, tag string) bool) {
	iterateTags(m.EncodedTags, int(host.TagIndex), cb)
}

func (m *CollectorConnections) GetContainerTags(container *ContainerMetadata) []string {
	return m.GetTags(int(container.TagIndex))
}

func (m *CollectorConnections) IterateContainerTags(container *ContainerMetadata, cb func(i, total int, tag string) bool) {
	iterateTags(m.EncodedTags, int(container.TagIndex), cb)
}

func (m *CollectorConnections) GetTags(tagIndex int) []string {
	return getTags(m.EncodedTags, tagIndex)
}

// GetDNS returns the DNS entries for the given addr.
// The first argument returned is the first DNS entry followed by any additional name resolutions.  Most IPs will
// have a single resolution so this dual format allows us to avoid allocations for the common case.  If there are
// multiple name resolutions, there is no implied priority between the dual values
func (m *CollectorConnections) GetDNS(addr *Addr) (string, []string) {
	return GetDNS(m.EncodedDNS, addr.Ip)
}

// IterateDNS iterates over all of the DNS entries for the given addr, invoking the provided callback for each one
func (m *CollectorConnections) IterateDNS(addr *Addr, cb func(i, total int, entry string) bool) {
	IterateDNS(m.EncodedDNS, addr.Ip, cb)
}

// GetDNSNames returns all the DNS entries
func (m *CollectorConnections) GetDNSNames() []string {
	return getDNSNames(m.EncodedDNS)
}
