package process

import (
	"fmt"
	"runtime"
	"testing"
)

func generateExtremeDuplicationTags() [][]string {
	uniqueTags := []string{"common", "shared", "reused", "frequent", "popular"}
	var tagGroups [][]string
	for i := 0; i < 10; i++ {
		set := make([]string, 50)
		for j := 0; j < 50; j++ {
			set[j] = uniqueTags[j%5]
		}
		tagGroups = append(tagGroups, set)
	}
	return tagGroups
}

func generateMicroservicesRealisticTags() [][]string {
	return [][]string{
		{"service:web", "env:prod", "datacenter:us-east", "team:platform", "version:1.2.3"},
		{"service:api", "env:prod", "datacenter:us-east", "team:platform", "version:1.2.3"},
		{"service:db", "env:prod", "datacenter:us-east", "team:data", "version:1.2.3"},
		{"service:cache", "env:prod", "datacenter:us-east", "team:platform", "version:1.2.3"},
		{"service:queue", "env:prod", "datacenter:us-east", "team:messaging", "version:1.2.3"},
		{"service:web", "env:staging", "datacenter:us-east", "team:platform", "version:1.2.4"},
		{"service:api", "env:staging", "datacenter:us-east", "team:platform", "version:1.2.4"},
		{"service:db", "env:staging", "datacenter:us-east", "team:data", "version:1.2.4"},
	}
}

func generateXLarge100Tags() [][]string {
	tags := make([]string, 100)
	for i := 0; i < 100; i++ {
		tags[i] = fmt.Sprintf("tag%03d:value%03d", i, i)
	}
	return [][]string{tags}
}

func BenchmarkTagEncodersAndDecoders(b *testing.B) {
	files := []struct {
		name  string
		files []string
		tags  [][]string
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
		{
			name: "extreme_duplication",
			tags: generateExtremeDuplicationTags(),
		},
		{
			name: "microservices_realistic",
			tags: generateMicroservicesRealisticTags(),
		},
		{
			name: "xlarge_100_tags",
			tags: generateXLarge100Tags(),
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
			name:           "v3",
			encoderFactory: NewV3TagEncoder,
		},
		{
			name:           "v1",
			encoderFactory: NewTagEncoder,
		},
	}

	for _, tt := range files {
		var tagGroups [][]string

		if tt.tags != nil {
			tagGroups = tt.tags
		} else {
			for _, file := range tt.files {
				tagGroups = append(tagGroups, readTestTags(b, file)...)
			}
		}

		for _, e := range encoders {
			// ENCODING BENCHMARK
			name := fmt.Sprintf("encode_%s_%s", tt.name, e.name)
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

			// DECODING BENCHMARK
			// Pre-encode for decoding test
			encoder := encoderFactory()
			var indices []int
			for _, tags := range tagGroups {
				idx := encoder.Encode(tags)
				indices = append(indices, idx)
			}
			buffer := encoder.Buffer()

			name = fmt.Sprintf("decode_%s_%s", tt.name, e.name)
			b.Run(name, func(b *testing.B) {
				b.ReportAllocs()

				var decoded []string
				for i := 0; i < b.N; i++ {
					for _, idx := range indices {
						decoded = getTags(buffer, idx)
					}
				}
				b.ReportMetric(float64(len(buffer)), "bytes")
				runtime.KeepAlive(decoded)
			})
		}
	}
}
