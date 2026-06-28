package process

import (
	"bytes"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNetworkPathTestIdentityWireEncoding(t *testing.T) {
	const testIdentity = "np-test-identity"

	networkPathBytes := []byte{0x08, 0x01, 0x12, byte(len(testIdentity))}
	networkPathBytes = append(networkPathBytes, testIdentity...)

	t.Run("network path message", func(t *testing.T) {
		encoded, err := proto.Marshal(&NetworkPath{
			HasTest:      true,
			TestIdentity: testIdentity,
		})
		require.NoError(t, err)
		assert.Equal(t, networkPathBytes, encoded)

		var decoded NetworkPath
		require.NoError(t, proto.Unmarshal(encoded, &decoded))
		assert.True(t, decoded.GetHasTest())
		assert.Equal(t, testIdentity, decoded.GetTestIdentity())
	})

	t.Run("connection field number", func(t *testing.T) {
		expected := []byte{0x8a, 0x04, byte(len(networkPathBytes))}
		expected = append(expected, networkPathBytes...)

		encoded, err := proto.Marshal(&Connection{
			NetworkPath: &NetworkPath{
				HasTest:      true,
				TestIdentity: testIdentity,
			},
		})
		require.NoError(t, err)
		assert.Equal(t, expected, encoded)

		var decoded Connection
		require.NoError(t, proto.Unmarshal(encoded, &decoded))
		require.NotNil(t, decoded.GetNetworkPath())
		assert.True(t, decoded.GetNetworkPath().GetHasTest())
		assert.Equal(t, testIdentity, decoded.GetNetworkPath().GetTestIdentity())
	})

	t.Run("empty network path is explicit false without identity", func(t *testing.T) {
		encoded, err := proto.Marshal(&Connection{NetworkPath: &NetworkPath{}})
		require.NoError(t, err)
		assert.Equal(t, []byte{0x8a, 0x04, 0x00}, encoded)

		var decoded Connection
		require.NoError(t, proto.Unmarshal(encoded, &decoded))
		require.NotNil(t, decoded.GetNetworkPath())
		assert.False(t, decoded.GetNetworkPath().GetHasTest())
		assert.Empty(t, decoded.GetNetworkPath().GetTestIdentity())
	})
}

func TestNetworkPathTestIdentityCollectorConnectionsRoundTrip(t *testing.T) {
	const testIdentity = "sha256_128-test-identity"

	payload := &CollectorConnections{
		HostName: "agent-hostname",
		Connections: []*Connection{
			{
				Pid: 101,
				NetworkPath: &NetworkPath{
					HasTest:      true,
					TestIdentity: testIdentity,
				},
			},
			{
				Pid:         202,
				NetworkPath: &NetworkPath{},
			},
			{
				Pid: 303,
			},
		},
	}

	encoded, err := proto.Marshal(payload)
	require.NoError(t, err)

	var decoded CollectorConnections
	require.NoError(t, proto.Unmarshal(encoded, &decoded))
	require.Len(t, decoded.GetConnections(), 3)

	withTest := decoded.GetConnections()[0].GetNetworkPath()
	require.NotNil(t, withTest)
	assert.True(t, withTest.GetHasTest())
	assert.Equal(t, testIdentity, withTest.GetTestIdentity())

	explicitFalse := decoded.GetConnections()[1].GetNetworkPath()
	require.NotNil(t, explicitFalse)
	assert.False(t, explicitFalse.GetHasTest())
	assert.Empty(t, explicitFalse.GetTestIdentity())

	assert.Nil(t, decoded.GetConnections()[2].GetNetworkPath())
}

func TestNetworkPathTestIdentityBuilderEncoding(t *testing.T) {
	const testIdentity = "builder-test-identity"

	expected, err := proto.Marshal(&Connection{
		NetworkPath: &NetworkPath{
			HasTest:      true,
			TestIdentity: testIdentity,
		},
	})
	require.NoError(t, err)

	var buf bytes.Buffer
	builder := NewConnectionBuilder(&buf)
	builder.SetNetworkPath(func(w *NetworkPathBuilder) {
		w.SetHasTest(true)
		w.SetTestIdentity(testIdentity)
	})
	assert.Equal(t, expected, buf.Bytes())

	var decoded Connection
	require.NoError(t, proto.Unmarshal(buf.Bytes(), &decoded))
	require.NotNil(t, decoded.GetNetworkPath())
	assert.True(t, decoded.GetNetworkPath().GetHasTest())
	assert.Equal(t, testIdentity, decoded.GetNetworkPath().GetTestIdentity())

	buf.Reset()
	networkPathBuilder := NewNetworkPathBuilder(&buf)
	networkPathBuilder.SetHasTest(true)
	networkPathBuilder.SetTestIdentity("")
	assert.Equal(t, []byte{0x08, 0x01}, buf.Bytes())
}

func TestNetworkPathTestIdentityJSONEncoding(t *testing.T) {
	const testIdentity = "json-test-identity"

	payload := &CollectorConnections{
		Connections: []*Connection{
			{
				NetworkPath: &NetworkPath{
					HasTest:      true,
					TestIdentity: testIdentity,
				},
			},
		},
	}

	var buf bytes.Buffer
	require.NoError(t, (&jsonpb.Marshaler{}).Marshal(&buf, payload))

	encoded := buf.String()
	assert.Contains(t, encoded, `"networkPath":`)
	assert.Contains(t, encoded, `"hasTest":true`)
	assert.Contains(t, encoded, `"testIdentity":"`+testIdentity+`"`)
}
