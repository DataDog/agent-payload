package process

import (
	"testing"

	"github.com/stretchr/testify/require"
	"inet.af/netaddr"
)

func TestIPMarshalJSON(t *testing.T) {
	orig := "10.1.2.3"
	ip := IP{netaddr.MustParseIP(orig)}
	b, err := ip.MarshalJSON()
	require.NoError(t, err)

	var ip2 IP
	err = ip2.UnmarshalJSON(b)
	require.NoError(t, err)
	require.Equal(t, ip, ip2)
}
