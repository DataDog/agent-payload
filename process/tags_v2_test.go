package process

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestV2TagEncoder(t *testing.T) {
	suite.Run(t, &TagSerdeTestSuite{encoder: NewV2TagEncoder()})
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
			name: "v2_no_map_pool_map_size_100",
			encoderFactory: func() TagEncoder {
				return NewV2TagEncoderWithOptions(false, 100)
			},
		},
		{
			name: "v2_no_map_pool_map_size_200",
			encoderFactory: func() TagEncoder {
				return NewV2TagEncoderWithOptions(false, 200)
			},
		},
		{
			name: "v2_no_map_pool",
			encoderFactory: func() TagEncoder {
				return NewV2TagEncoderWithOptions(false, 0)
			},
		},
		{
			name:           "v2_map_pool",
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
				runtime.KeepAlive(buf)
			})
		}
	}
}

func BenchmarkMaps(b *testing.B) {
	files := []string{
		"testdata/low_dups.txt",
		"testdata/high_dups.txt",
		"testdata/high_dups_2.txt",
	}

	for _, file := range files {
		tagGroups := readTestTags(b, file)

		name := fmt.Sprintf("intkey_%s", filepath.Base(file[:strings.Index(file, ".")]))
		b.Run(name, func(b *testing.B) {
			b.ReportAllocs()

			var data map[string]uint32
			for i := 0; i < b.N; i++ {
				data = make(map[string]uint32)
				for i, tags := range tagGroups {
					for j, tag := range tags {
						data[tag] = uint32(i + j)
					}
				}

			}
			runtime.KeepAlive(data)
		})

		name = fmt.Sprintf("emptykey_%s", filepath.Base(file[:strings.Index(file, ".")]))
		b.Run(name, func(b *testing.B) {
			b.ReportAllocs()

			var data map[string]struct{}
			for i := 0; i < b.N; i++ {
				data = make(map[string]struct{})
				for _, tags := range tagGroups {
					for _, tag := range tags {
						data[tag] = struct{}{}
					}
				}
			}
			runtime.KeepAlive(data)
		})
	}
}
