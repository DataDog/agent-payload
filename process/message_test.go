package process

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestDecodeZstd05Payload ensures backward compatibility with our intake
func TestDecodeZstd05Payload(t *testing.T) {
	file := "testdata/test_zstd.0.5.dump"
	expected := Message{
		Header: MessageHeader{
			Version:  MessageV3,
			Encoding: MessageEncodingZstdPB,
			Type:     TypeCollectorProc,
		},
		Body: &CollectorProc{
			HostName: "test",
		},
	}

	raw, err := ioutil.ReadFile(file)
	assert.NoError(t, err)

	msg, err := DecodeMessage(raw)
	assert.NoError(t, err)

	assert.Equal(t, expected, msg)
}

func TestMessageTypeString(t *testing.T) {
	cases := map[MessageType]string{
		TypeCollectorProc:              "process",
		TypeCollectorConnections:       "network",
		TypeCollectorRealTime:          "process-rt",
		TypeCollectorContainer:         "container",
		TypeCollectorContainerRealTime: "container-rt",
		TypeCollectorPod:               "pod",
		TypeCollectorReplicaSet:        "replica-set",
		TypeCollectorDeployment:        "deployment",
		TypeCollectorService:           "service",
		TypeResCollector:               "23",
	}
	for input, expected := range cases {
		assert.Equal(t, input.String(), expected)
	}
}
