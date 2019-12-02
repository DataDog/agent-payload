package process

import (
	"encoding/binary"
	"math"
	"sync"
)

// 1 meta byte + 4 bytes for the index of the footer block
const v2PreambleLength = 1 + 4

type V2TagEncoder struct {
	tags          map[string]uint32
	order         []string
	footer        []byte
	tagPosition   uint32
	usePooledMaps bool
}

var footerPool = sync.Pool{
	New: func() interface{} {
		var footer []byte
		return &footer
	},
}

var orderPool = sync.Pool{
	New: func() interface{} {
		var order []string
		return &order
	},
}

var tagsPool = sync.Pool{
	New: func() interface{} {
		tags := make(map[string]uint32)
		return &tags
	},
}

func NewV2TagEncoder() TagEncoder {
	return NewV2TagEncoderWithOptions(true, 0)
}

func NewV2TagEncoderWithOptions(usePooledMaps bool, mapSize int) TagEncoder {
	footer := *footerPool.Get().(*[]byte)
	order := *orderPool.Get().(*[]string)

	var tags map[string]uint32
	if usePooledMaps {
		tags = *tagsPool.Get().(*map[string]uint32)
	} else {
		tags = make(map[string]uint32, mapSize)
	}

	return &V2TagEncoder{
		tags:          tags,
		order:         order[:0],
		footer:        footer[:0],
		tagPosition:   v2PreambleLength,
		usePooledMaps: usePooledMaps,
	}
}

func (t *V2TagEncoder) Buffer() []byte {
	tagsSize := 0

	for _, tag := range t.order {
		tagsSize += 2 + len(tag)
	}

	bufferSize := v2PreambleLength + tagsSize + len(t.footer)

	footerPosition := uint32(v2PreambleLength + tagsSize)

	buffer := make([]byte, 0, bufferSize)
	buffer = append(buffer, version2)

	var intBuf [4]byte
	binary.LittleEndian.PutUint32(intBuf[:], footerPosition)
	buffer = append(buffer, intBuf[:]...)

	var shortBuf [2]byte
	for _, tag := range t.order {
		binary.LittleEndian.PutUint16(shortBuf[:], uint16(len(tag)))
		buffer = append(buffer, shortBuf[:]...)
		buffer = append(buffer, tag...)
	}

	buffer = append(buffer, t.footer...)

	footerPool.Put(&t.footer)
	orderPool.Put(&t.order)

	if t.usePooledMaps {
		for k := range t.tags {
			delete(t.tags, k)
		}

		tagsPool.Put(&t.tags)
	}

	return buffer
}

func (t *V2TagEncoder) Encode(tags []string) int {
	if len(tags) == 0 {
		return -1
	}

	// We only allow 2 bytes for the number of the tags, ensure we don't exceed it
	if len(tags) > math.MaxUint16 {
		tags = tags[0:math.MaxUint16]
	}

	// The index for these tags is the current end of the footer
	tagIndex := len(t.footer)

	var shortBuf [2]byte
	var intBuf [4]byte

	binary.LittleEndian.PutUint16(shortBuf[:], uint16(len(tags)))
	t.footer = append(t.footer, shortBuf[:]...)

	for _, tag := range tags {
		// We only allow 2 bytes for the length of the tag, ensure we don't exceed it
		if len(tag) > math.MaxUint16 {
			tag = tag[0:math.MaxUint16]
		}

		position, ok := t.tags[tag]
		if !ok {
			position = t.tagPosition
			t.tagPosition += uint32(2 + len(tag))
			t.tags[tag] = position
			t.order = append(t.order, tag)
		}

		binary.LittleEndian.PutUint32(intBuf[:], position)
		t.footer = append(t.footer, intBuf[:]...)
	}

	return tagIndex
}

func decodeV2(buffer []byte, tagIndex int) []string {
	footerPosition := binary.LittleEndian.Uint32(buffer[1:])

	idx := int(footerPosition) + tagIndex

	footerBuffer := buffer[idx:]
	footerIndex := 0

	numTags := int(binary.LittleEndian.Uint16(footerBuffer[footerIndex:]))
	footerIndex += 2

	tags := make([]string, 0, numTags)

	for i := 0; i < numTags; i++ {
		tagPosition := int(binary.LittleEndian.Uint32(footerBuffer[footerIndex:]))

		tagLength := int(binary.LittleEndian.Uint16(buffer[tagPosition:]))

		tag := string(buffer[tagPosition+2 : tagPosition+2+tagLength])
		tags = append(tags, tag)

		footerIndex += 4
	}

	return tags
}
