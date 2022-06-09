package process

func (m IPAddress) IsZero() bool {
	return m.Low == 0 && m.High == 0
}
