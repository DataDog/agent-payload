package process

type ContainerAddrNoProto struct {
	Ip       string
	Port     int32
	Protocol ConnectionType
}

func (c *ContainerAddr) ToNoProto() ContainerAddrNoProto {
	return ContainerAddrNoProto{
		Ip:       c.Ip,
		Port:     c.Port,
		Protocol: c.Protocol,
	}
}
