//go:build !cgo

package process

import (
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

func TestManifestPayloadAllEncodings_NoCGO(t *testing.T) {
	testCases := []struct {
		name     string
		encoding MessageEncoding
	}{
		{"Protobuf", MessageEncodingProtobuf},
		{"JSON", MessageEncodingJSON},
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

func TestUnsupportedEncodings_NoCGO(t *testing.T) {
	unsupportedEncodings := []struct {
		name     string
		encoding MessageEncoding
	}{
		{"ZstdPB", MessageEncodingZstdPB},
		{"Zstd1xPB", MessageEncodingZstd1xPB},
	}

	for _, tc := range unsupportedEncodings {
		t.Run(tc.name+"_EncodingError", func(t *testing.T) {
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

			// These should fail in non-CGO builds during encoding
			_, err := EncodeMessage(message)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), "unsupported encoding",
				"Expected unsupported encoding error for %s", tc.name)
		})

		t.Run(tc.name+"_DecodingError", func(t *testing.T) {
			// Test decoding error by creating a minimal message with unsupported encoding
			header := MessageHeader{
				Version:  MessageV3,
				Encoding: tc.encoding,
				Type:     TypeCollectorManifest,
			}

			// Create a basic protobuf-encoded body
			body := &CollectorManifest{HostName: "test"}
			bodyBytes, err := proto.Marshal(body)
			assert.NoError(t, err)

			// Encode header
			headerBytes, err := encodeHeader(header)
			assert.NoError(t, err)

			// Combine header and body to simulate a message with unsupported encoding
			fakeMessage := append(headerBytes, bodyBytes...)

			// This should fail during decoding
			_, err = DecodeMessage(fakeMessage)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), "unsupported encoding",
				"Expected unsupported encoding error when decoding %s", tc.name)
		})
	}
}
