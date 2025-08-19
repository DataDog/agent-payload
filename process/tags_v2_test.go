package process

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestV2TagEncoder(t *testing.T) {
	suite.Run(t, &TagSerdeTestSuite{encoder: NewV2TagEncoder()})
}

func TestUnsafeIterationV2(t *testing.T) {
	// buff = 2
	buf := make([]byte, 2)
	assert.NotPanics(t, func() {
		unsafeIterateV2(buf, 0, func(i, total int, tag []byte) bool { return true })
	})

	// buff = 1
	assert.NotPanics(t, func() {
		unsafeIterateV2(buf[1:], 0, func(i, total int, tag []byte) bool { return true })
	})

	// indx > buff
	buf = make([]byte, 6)
	assert.NotPanics(t, func() {
		unsafeIterateV2(buf, 10, func(i, total int, tag []byte) bool { return true })
	})

	// footerBuffer < 2
	assert.NotPanics(t, func() {
		unsafeIterateV2(buf, 5, func(i, total int, tag []byte) bool { return true })
	})
}

func FuzzIterateV2(f *testing.F) {
	allTags := readTestTags(f, "testdata/tags.txt")
	t := NewV2TagEncoder()

	for _, tag := range allTags {
		_ = t.Encode(tag)
	}
	buf := t.Buffer()

	f.Add(buf, 1)
	f.Fuzz(FuzzingIterateV2)
}

func FuzzingIterateV2(t *testing.T, buffer []byte, tagIndex int) {
	assert.NotPanics(t, func() {
		unsafeIterateV2(buffer, tagIndex, func(i, total int, tag []byte) bool { return true })
	})

}
