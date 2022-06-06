package process

func (m *Addr) GetIP() IPAddress {
	if m.IpAddr == nil {
		return IPAddress{}
	}
	return *m.IpAddr
}

func (m *ContainerAddr) GetIP() IPAddress {
	if m.IpAddr == nil {
		return IPAddress{}
	}
	return *m.IpAddr
}

func (m IPAddress) IsZero() bool {
	return m.Low == 0 && m.High == 0
}
