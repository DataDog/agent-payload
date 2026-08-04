//go:build cgo

package process

import (
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

func TestMessageTypeString(t *testing.T) {
	cases := map[MessageType]string{
		TypeCollectorProc:                  "process",
		TypeCollectorConnections:           "network",
		TypeCollectorRealTime:              "process-rt",
		TypeCollectorContainer:             "container",
		TypeCollectorContainerRealTime:     "container-rt",
		TypeCollectorPod:                   "pod",
		TypeCollectorReplicaSet:            "replica-set",
		TypeCollectorDeployment:            "deployment",
		TypeCollectorService:               "service",
		TypeCollectorNode:                  "node",
		TypeCollectorCluster:               "cluster",
		TypeCollectorManifest:              "manifest",
		TypeCollectorJob:                   "job",
		TypeCollectorCronJob:               "cron-job",
		TypeCollectorDaemonSet:             "daemon-set",
		TypeCollectorStatefulSet:           "stateful-set",
		TypeCollectorPersistentVolume:      "persistent-volume",
		TypeCollectorPersistentVolumeClaim: "persistent-volume-claim",
		TypeCollectorProcDiscovery:         "process-discovery",
		TypeCollectorRole:                  "role",
		TypeCollectorRoleBinding:           "role-binding",
		TypeCollectorClusterRole:           "cluster-role",
		TypeCollectorClusterRoleBinding:    "cluster-role-binding",
		TypeCollectorServiceAccount:        "service-account",
		TypeCollectorIngress:               "ingress",
		TypeCollectorProcEvent:             "process-event",
		TypeResCollector:                   "23",
		TypeCollectorNetworkPolicy:         "network-policy",
		TypeCollectorPodDisruptionBudget:   "pod-disruption-budget",
	}
	for input, expected := range cases {
		assert.Equal(t, expected, input.String())
	}
}

func TestManifestPayloadAllEncodings_CGO(t *testing.T) {
	testCases := []struct {
		name     string
		encoding MessageEncoding
	}{
		{"Protobuf", MessageEncodingProtobuf},
		{"JSON", MessageEncodingJSON},
		{"Zstd1xPB", MessageEncodingZstd1xPB},
		{"ZstdPBxNoCgo", MessageEncodingZstdPBxNoCgo},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			message := Message{
				Header: MessageHeader{
					Version:  MessageV3,
					Encoding: tc.encoding,
					Type:     TypeCollectorManifest,
				},
				Body: &CollectorManifest{
					HostName: "test",
				},
			}

			encoded, err := EncodeMessage(message)
			assert.NoError(t, err)

			msg, err := DecodeMessage(encoded)
			assert.NoError(t, err)

			// For JSON encoding, compare the essential fields since JSON marshaling
			// can change nil slices to empty slices
			if tc.encoding == MessageEncodingJSON {
				assert.Equal(t, message.Header, msg.Header)
				originalBody := message.Body.(*CollectorManifest)
				decodedBody := msg.Body.(*CollectorManifest)
				assert.Equal(t, originalBody.HostName, decodedBody.HostName)
				assert.Equal(t, originalBody.GroupSize, decodedBody.GroupSize)
			} else {
				assert.Equal(t, message, msg)
			}
		})
	}
}

// TestRemovedZstdPBEncoding ensures the retired MessageEncodingZstdPB wire
// value (2, formerly backed by the abandoned zstd_0 module) is rejected as
// an unknown encoding rather than silently accepted.
func TestRemovedZstdPBEncoding(t *testing.T) {
	const removedZstdPBEncoding MessageEncoding = 2

	message := Message{
		Header: MessageHeader{
			Version:  MessageV3,
			Encoding: removedZstdPBEncoding,
			Type:     TypeCollectorManifest,
		},
		Body: &CollectorManifest{
			HostName: "test",
		},
	}

	t.Run("EncodingError", func(t *testing.T) {
		_, err := EncodeMessage(message)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "unknown message encoding")
	})

	t.Run("DecodingError", func(t *testing.T) {
		bodyBytes, err := proto.Marshal(message.Body)
		assert.NoError(t, err)

		headerBytes, err := encodeHeader(message.Header)
		assert.NoError(t, err)

		fakeMessage := append(headerBytes, bodyBytes...)

		_, err = DecodeMessage(fakeMessage)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "unknown message encoding")
	})
}
