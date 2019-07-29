package process

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTagSerde(t *testing.T) {
	encoder := NewTagEncoder()

	a := encoder.Encode([]string{"one", "two", "three"})
	b := encoder.Encode([]string{})
	c := encoder.Encode([]string{"four", "five"})
	d := encoder.Encode([]string{"six"})
	e := encoder.Encode([]string{"seven", "eight", "nine", "ten"})

	assert.Equal(t, []string{"one", "two", "three"}, getTags(encoder.Buffer(), a))
	assert.Equal(t, []string{}, getTags(encoder.Buffer(), b))
	assert.Equal(t, []string{"four", "five"}, getTags(encoder.Buffer(), c))
	assert.Equal(t, []string{"six"}, getTags(encoder.Buffer(), d))
	assert.Equal(t, []string{"seven", "eight", "nine", "ten"}, getTags(encoder.Buffer(), e))
}

func TestTagSerdeRealTags(t *testing.T) {
	allTags := readTestTags(t)

	encoder := NewTagEncoder()

	var tagIndices []int

	for _, tags := range allTags {
		tagIndex := encoder.Encode(tags)
		tagIndices = append(tagIndices, tagIndex)
	}

	for i, tagIndex := range tagIndices {
		assert.Equal(t, allTags[i], getTags(encoder.Buffer(), tagIndex))
	}
}

func TestGetTagsEmpty(t *testing.T) {
	assert.Empty(t, getTags(nil, 1234))
}

func TestOverflowNumberOfTags(t *testing.T) {
	var tags []string

	for i := 0; i < math.MaxUint16+1; i++ {
		tags = append(tags, fmt.Sprintf("%d", i))
	}

	encoder := NewTagEncoder()

	idx := encoder.Encode(tags)

	assert.Len(t, getTags(encoder.Buffer(), idx), math.MaxUint16)
}

func TestOverflowTagLength(t *testing.T) {
	tag := ""

	for i := 0; i < math.MaxUint16+1; i++ {
		tag += "0"
	}

	encoder := NewTagEncoder()

	idx := encoder.Encode([]string{tag})

	tags := getTags(encoder.Buffer(), idx)

	require.Len(t, tags, 1)
	assert.Len(t, tags[0], math.MaxUint16)
}

func BenchmarkTagEncode(b *testing.B) {
	allTags := readTestTags(b)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t := NewTagEncoder()

		for _, tags := range allTags {
			_ = t.Encode(tags)
		}
	}
}

func readTestTags(t require.TestingT) [][]string {
	buf, err := ioutil.ReadFile("testdata/tags.txt")
	require.NoError(t, err)

	var allTags [][]string
	for _, line := range strings.Split(string(buf), "\n") {
		tags := strings.Split(line, " ")
		allTags = append(allTags, tags)
	}

	return allTags
}
