package process

import (
	"encoding/binary"
	"math"
)

type TagEncoder interface {
	// Buffer returns the underlying byte buffer that the tags were encoded in to
	Buffer() []byte

	// Encode encodes the given tags in to the buffer and returns the index in the buffer where the data begins
	Encode(tags []string) int
}

// Version for the encoding format
const version1 = 1

// Groups of tags are successively encoded in to a single buffer. For each group of tags, the format is:
// - Number of tags encoded as a 2-byte uint16.
// - For each tag, write the length of the tag as a 2-byte uint16 followed by the tag bytes.
type v1TagEncoder struct {
	buffer []byte
}

// NewTagEncoder creates an empty tag encoder
func NewTagEncoder() TagEncoder {
	// Reserve the first byte to version the format
	initialBuf := []byte{version1}

	return &v1TagEncoder{buffer: initialBuf}
}

func (t *v1TagEncoder) Buffer() []byte {
	return t.buffer
}

func (t *v1TagEncoder) Encode(tags []string) int {
	// We only allow 2 bytes for the number of the tags, ensure we don't exceed it
	if len(tags) > math.MaxUint16 {
		tags = tags[0:math.MaxUint16]
	}

	bufferSize := bufferSize(tags)

	// Check to see if there is enough space in the buffer that we can reuse rather than allocating a temporary buffer
	newBufferRequired := (cap(t.buffer) - len(t.buffer)) < bufferSize

	tagBuffer := t.buffer[len(t.buffer):]

	if newBufferRequired {
		tagBuffer = make([]byte, 0, bufferSize)
	}

	var sizeBuf [2]byte
	binary.LittleEndian.PutUint16(sizeBuf[0:], uint16(len(tags)))
	tagBuffer = append(tagBuffer, sizeBuf[0:]...)

	for _, tag := range tags {
		// We only allow 2 bytes for the length of the tag, ensure we don't exceed it
		if len(tag) > math.MaxUint16 {
			tag = tag[0:math.MaxUint16]
		}

		binary.LittleEndian.PutUint16(sizeBuf[0:], uint16(len(tag)))
		tagBuffer = append(tagBuffer, sizeBuf[0:]...)
		tagBuffer = append(tagBuffer, tag...)
	}

	// The index for these tags is the current end of the buffer
	tagIndex := len(t.buffer)

	if newBufferRequired {
		t.buffer = append(t.buffer, tagBuffer...)
	} else {
		t.buffer = t.buffer[0 : len(t.buffer)+bufferSize]
	}

	return tagIndex
}

func getTags(buffer []byte, tagIndex int) []string {
	if len(buffer) == 0 {
		return nil
	}

	tagBuffer := buffer[tagIndex:]
	readIndex := 0

	numTags := int(binary.LittleEndian.Uint16(tagBuffer[readIndex:]))
	readIndex += 2

	tags := make([]string, 0, numTags)

	for i := 0; i < numTags; i++ {
		tagLength := int(binary.LittleEndian.Uint16(tagBuffer[readIndex:]))
		readIndex += 2

		tag := string(tagBuffer[readIndex : readIndex+tagLength])
		tags = append(tags, tag)

		readIndex += tagLength
	}

	return tags
}

// bufferSize returns the number of bytes required to store the given tags
func bufferSize(tags []string) int {
	// Include space for the number of tags
	bufferSize := 2

	for _, tag := range tags {
		// Include space for the length of the tag and the tag itself
		bufferSize += 2 + len(tag)
	}

	return bufferSize
}
