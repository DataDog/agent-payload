package process

import (
	"encoding/json"
	"inet.af/netaddr"
)

type IP struct {
	netaddr.IP
}

func (t IP) Marshal() ([]byte, error) {
	return t.MarshalBinary()
}

func (t *IP) MarshalTo(data []byte) (n int, err error) {
	if t.IsZero() {
		return
	}
	// TODO these could be improved without the intermediate array
	// if we had access to the underlying uint128
	if t.Is4() {
		b := t.IP.As4()
		n = copy(data, b[:])
		return
	}
	b := t.IP.As16()
	n = copy(data, b[:])
	if z := t.IP.Zone(); z != "" {
		n += copy(data[n:], z)
	}
	return
}

func (t *IP) Unmarshal(data []byte) error {
	return t.UnmarshalBinary(data)
}

func (t *IP) Size() int {
	return int(t.BitLen() / 8)
}

func (t IP) MarshalJSON() ([]byte, error) {
	b, err := t.IP.MarshalText()
	if err != nil {
		return nil, err
	}
	return json.Marshal(b)
}

func (t *IP) UnmarshalJSON(data []byte) error {
	var ip string
	err := json.Unmarshal(data, &ip)
	if err != nil {
		return err
	}
	return t.IP.UnmarshalText([]byte(ip))
}
