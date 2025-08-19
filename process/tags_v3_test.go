package process

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestV3TagEncoder(t *testing.T) {
	suite.Run(t, &TagSerdeTestSuite{encoder: NewV3TagEncoder(), shouldApplySorting: true, shouldApplyDedup: true})
}

func TestUnsafeIterationV3(t *testing.T) {
	// buff = 3
	buf := make([]byte, 3)
	assert.NotPanics(t, func() {
		unsafeIterateV3(buf, 0, func(i, total int, tag []byte) bool { return true })
	})

	// buff = 1
	assert.NotPanics(t, func() {
		unsafeIterateV3(buf[1:], 0, func(i, total int, tag []byte) bool { return true })
	})

	// indx > buff
	buf = make([]byte, 6)
	assert.NotPanics(t, func() {
		unsafeIterateV3(buf, 10, func(i, total int, tag []byte) bool { return true })
	})

	// footerBuffer < 2
	assert.NotPanics(t, func() {
		unsafeIterateV3(buf, 5, func(i, total int, tag []byte) bool { return true })
	})
}

func FuzzIterateV3(f *testing.F) {
	allTags := readTestTags(f, "testdata/tags.txt")
	t := NewV3TagEncoder()

	for _, tag := range allTags {
		_ = t.Encode(tag)
	}
	buf := t.Buffer()

	f.Add(buf, 1)
	f.Fuzz(FuzzingIterateV3)
}

func FuzzingIterateV3(t *testing.T, buffer []byte, tagIndex int) {
	assert.NotPanics(t, func() {
		unsafeIterateV3(buffer, tagIndex, func(i, total int, tag []byte) bool { return true })
	})

}
