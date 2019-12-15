package process

func (m *CollectorConnections) GetHostTags(host *Host) []string {
	return m.GetTags(int(host.TagIndex))
}

func (m *CollectorConnections) GetContainerTags(container *ContainerMetadata) []string {
	return m.GetTags(int(container.TagIndex))
}

func (m *CollectorConnections) GetTags(tagIndex int) []string {
	return getTags(m.EncodedTags, tagIndex)
}

// GetDNS returns the DNS entries for the given addr.
// The first argument returned is the first DNS entry followed by any additional name resolutions.  Most IPs will
// have a single resolution so this dual format allows us to avoid allocations for the common case.  If there are
// multiple name resolutions, there is no implied priority between the dual values
func (m *CollectorConnections) GetDNS(addr *Addr) (string, []string) {
	return getDNS(m.EncodedDNS, addr.Ip)
}
