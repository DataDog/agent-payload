package process

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIterateDNS(t *testing.T) {
	t.Run("v1", func(t *testing.T) {
		addr := &Addr{Ip: "1.1.1.1"}
		encoder := NewV1DNSEncoder()
		buf, err := encoder.Encode(map[string]*DNSEntry{
			addr.Ip: {Names: []string{"foo", "bar"}},
		})
		require.NoError(t, err)

		cc := &CollectorConnections{
			EncodedDNS: buf,
		}

		var entries []string
		err = cc.IterateDNS(addr, func(i, total int, entry string) bool {
			entries = append(entries, entry)
			return true
		})
		require.NoError(t, err)
		assert.ElementsMatch(t, []string{"foo", "bar"}, entries)

		entries = nil
		err = cc.UnsafeIterateDNS(addr, func(i, total int, entry []byte) bool {
			entries = append(entries, string(entry))
			return true
		})
		require.NoError(t, err)
		assert.ElementsMatch(t, []string{"foo", "bar"}, entries)
	})

	t.Run("v2", func(t *testing.T) {
		addr := &Addr{Ip: "1.1.1.1"}
		db := []string{"foo", "bar"}
		encoder := NewV2DNSEncoder()
		dbBuf, indexToOffset, err := encoder.EncodeDomainDatabase(db)
		require.NoError(t, err)

		lookupBuf, err := encoder.EncodeMapped(map[string]*DNSDatabaseEntry{
			addr.Ip: {NameOffsets: []int32{indexOf("foo", db), indexOf("bar", db)}},
		}, indexToOffset)
		require.NoError(t, err)

		cc := CollectorConnections{EncodedDnsLookups: lookupBuf, EncodedDomainDatabase: dbBuf}

		var entries []string
		err = cc.IterateDNS(addr, func(i, total int, entry string) bool {
			entries = append(entries, entry)
			return true
		})
		require.NoError(t, err)
		assert.ElementsMatch(t, []string{"foo", "bar"}, entries)

		entries = nil
		err = cc.UnsafeIterateDNS(addr, func(i, total int, entry []byte) bool {
			entries = append(entries, string(entry))
			return true
		})
		require.NoError(t, err)
		assert.ElementsMatch(t, []string{"foo", "bar"}, entries)
	})
}
