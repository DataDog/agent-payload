package process

func (m *Addr) GetIP() IPAddress {
	if m.IpAddr == nil {
		return IPAddress{}
	}
	return *m.IpAddr
}
