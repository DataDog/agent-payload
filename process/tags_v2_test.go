package process

import (
	"fmt"
	"runtime"
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
	t := NewTagEncoder()

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

func BenchmarkTagEncoders(b *testing.B) {
	files := []struct {
		name  string
		files []string
	}{
		{
			name:  "low_dups",
			files: []string{"testdata/low_dups.txt"},
		},
		{
			name:  "high_dups",
			files: []string{"testdata/high_dups.txt"},
		},
		{
			name:  "high_dups_2",
			files: []string{"testdata/high_dups_2.txt"},
		},
		{
			name:  "combined",
			files: []string{"testdata/low_dups.txt", "testdata/high_dups.txt", "testdata/high_dups_2.txt"},
		},
	}

	encoders := []struct {
		name           string
		encoderFactory func() TagEncoder
	}{
		{
			name:           "v2",
			encoderFactory: NewV2TagEncoder,
		},
		{
			name:           "v1",
			encoderFactory: NewTagEncoder,
		},
	}

	for _, tt := range files {
		var tagGroups [][]string
		for _, file := range tt.files {
			tagGroups = append(tagGroups, readTestTags(b, file)...)
		}

		for _, e := range encoders {
			name := fmt.Sprintf("%s_%s", tt.name, e.name)
			encoderFactory := e.encoderFactory

			b.Run(name, func(b *testing.B) {
				b.ReportAllocs()

				var buf []byte
				for i := 0; i < b.N; i++ {
					encoder := encoderFactory()
					for _, tags := range tagGroups {
						encoder.Encode(tags)
					}

					buf = encoder.Buffer()
				}
				b.ReportMetric(float64(len(buf)), "bytes")
				runtime.KeepAlive(buf)
			})
		}
	}
}
