package manifest

import (
	"bytes"

	"github.com/DataDog/zstd_0"
	proto "github.com/gogo/protobuf/proto"
)

// EncodeMessage is to marshal manifest then compress it by zstd
func EncodeMessage(mp *ManifestPayload) ([]byte, error) {
	b := new(bytes.Buffer)
	var p []byte
	pb, err := proto.Marshal(mp)
	if err != nil {
		return nil, err
	}

	p, err = zstd_0.Compress(nil, pb)
	if err != nil {
		return nil, err
	}

	_, err = b.Write(p)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

// DecodeMessage decompress manifest by zstd then unmarshal it
func DecodeMessage(data []byte) (*ManifestPayload, error) {
	var m = &ManifestPayload{}

	d, err := zstd_0.Decompress(nil, data)
	if err != nil {
		return nil, err
	}

	proto.Unmarshal(d, m)

	return m, nil
}
