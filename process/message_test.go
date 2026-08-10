//go:build cgo

package process

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestDecodeZstd05Payload ensures MessageEncodingZstdPB payloads fail to
// decode now that the vulnerable zstd_0 dependency has been removed.
func TestDecodeZstd05Payload(t *testing.T) {
	file := "testdata/test_zstd.0.5.dump"

	raw, err := ioutil.ReadFile(file)
	assert.NoError(t, err)

	_, err = DecodeMessage(raw)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unsupported encoding")
}

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

// TestZstdPBUnsupported_CGO ensures MessageEncodingZstdPB is rejected now
// that the vulnerable zstd_0 dependency has been removed.
func TestZstdPBUnsupported_CGO(t *testing.T) {
	message := Message{
		Header: MessageHeader{
			Version:  MessageV3,
			Encoding: MessageEncodingZstdPB,
			Type:     TypeCollectorManifest,
		},
		Body: &CollectorManifest{
			HostName: "test",
		},
	}

	_, err := EncodeMessage(message)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unsupported encoding")
}
