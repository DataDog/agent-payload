package process

import (
	"encoding/binary"
	"encoding/json"
	
	"github.com/gogo/protobuf/proto"
	"inet.af/netaddr"
)

type IP struct {
	netaddr.IP
}

func (t *IP) protoType() *ProtoIP {
	b := t.As16()
	return &ProtoIP{
		Hi: binary.BigEndian.Uint64(b[:8]),
		Lo: binary.BigEndian.Uint64(b[8:]),
	}
}

func (t IP) Marshal() ([]byte, error) {
	return proto.Marshal(t.protoType())
}

func (t *IP) MarshalTo(data []byte) (n int, err error) {
	return t.protoType().MarshalTo(data)
}

func (t *IP) Unmarshal(data []byte) error {
	pr := &ProtoIP{}
	err := proto.Unmarshal(data, pr)
	if err != nil {
		return err
	}

	var a [16]byte
	binary.BigEndian.PutUint64(a[:8], pr.Hi)
	binary.BigEndian.PutUint64(a[8:], pr.Lo)
	t.IP = netaddr.IPFrom16(a)
	return nil
}

func (t *IP) Size() int {
	return t.protoType().Size()
}

func (t IP) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.IP.String())
}

func (t *IP) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	ip, err := netaddr.ParseIP(s)
	if err != nil {
		return err
	}
	t.IP = ip
	return nil
}
