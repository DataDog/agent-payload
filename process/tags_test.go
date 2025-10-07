package process

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type TagSerdeTestSuite struct {
	suite.Suite
	encoder            TagEncoder
	shouldApplySorting bool
	shouldApplyDedup   bool
}

func applyModifiers(tags []string, shouldApplySorting, shouldApplyDedup bool) []string {
	if len(tags) == 0 {
		return tags
	}
	if shouldApplyDedup {
		deduped := make(map[string]struct{}, len(tags))
		for _, tag := range tags {
			deduped[tag] = struct{}{}
		}
		tags = make([]string, 0, len(deduped))
		for tag := range deduped {
			tags = append(tags, tag)
		}
	}
	if shouldApplySorting {
		sort.Strings(tags)
	}
	return tags
}

func TestV1TagEncoder(t *testing.T) {
	suite.Run(t, &TagSerdeTestSuite{encoder: NewTagEncoder()})
}

func (suite *TagSerdeTestSuite) TestTagSerde() {
	a := suite.encoder.Encode([]string{"one", "two", "three"})
	b := suite.encoder.Encode([]string{})
	c := suite.encoder.Encode([]string{"four", "five"})
	d := suite.encoder.Encode([]string{"six"})
	e := suite.encoder.Encode([]string{"seven", "eight", "nine", "ten"})

	buf := suite.encoder.Buffer()

	cases := map[int][]string{
		a: {"one", "two", "three"},
		b: nil,
		c: {"four", "five"},
		d: {"six"},
		e: {"seven", "eight", "nine", "ten"},
	}
	for idx, tags := range cases {
		assert.Equal(suite.T(), applyModifiers(tags, suite.shouldApplySorting, suite.shouldApplyDedup), getTags(buf, idx), fmt.Sprintf("for index %d", idx))
	}
}

func (suite *TagSerdeTestSuite) TestUnicodeTags() {
	encoder := suite.encoder

	tags := []string{"データベース", "ロガー", "english", "ウェブホスト"}

	a := encoder.Encode(tags)

	assert.Equal(suite.T(), tags, getTags(encoder.Buffer(), a))
}

func (suite *TagSerdeTestSuite) TestTagSerdeRealTags() {
	allTags := readTestTags(suite.T(), "testdata/tags.txt")

	encoder := suite.encoder

	var tagIndices []int

	for _, tags := range allTags {
		tagIndex := encoder.Encode(tags)
		tagIndices = append(tagIndices, tagIndex)
	}

	for i, tagIndex := range tagIndices {
		buffer := encoder.Buffer()
		expected := applyModifiers(allTags[i], suite.shouldApplySorting, suite.shouldApplyDedup)
		assert.Equal(suite.T(), expected, getTags(buffer, tagIndex))

		var iterated []string
		iterateTags(buffer, tagIndex, func(i, total int, tag string) bool {
			iterated = append(iterated, tag)
			return true
		})
		assert.Equal(suite.T(), expected, iterated)

		var unsafeIterated []string
		unsafeIterateTags(buffer, tagIndex, func(i, total int, tag []byte) bool {
			unsafeIterated = append(unsafeIterated, string(tag))
			return true
		})
		assert.Equal(suite.T(), expected, unsafeIterated)

		iterated = nil
		iterateTags(buffer, tagIndex, func(i, total int, tag string) bool {
			if i == total-1 {
				return false
			}
			iterated = append(iterated, tag)
			return true
		})
		assert.Equal(suite.T(), expected[:len(expected)-1], iterated)

	}
}

func (suite *TagSerdeTestSuite) TestGetTagsEmpty() {
	assert.Empty(suite.T(), getTags(nil, 1234))
}

func (suite *TagSerdeTestSuite) TestOverflowNumberOfTags() {
	var tags []string

	for i := 0; i < math.MaxUint16+1; i++ {
		tags = append(tags, fmt.Sprintf("%d", i))
	}

	idx := suite.encoder.Encode(tags)

	assert.Len(suite.T(), getTags(suite.encoder.Buffer(), idx), math.MaxUint16)
}

func (suite *TagSerdeTestSuite) TestOverflowTagLength() {
	tag := ""

	for i := 0; i < math.MaxUint16+1; i++ {
		tag += "0"
	}

	idx := suite.encoder.Encode([]string{tag})

	buffer := suite.encoder.Buffer()

	tags := getTags(buffer, idx)

	require.Len(suite.T(), tags, 1)
	assert.Len(suite.T(), tags[0], math.MaxUint16)
}

func TestV1DecodedTags(t *testing.T) {
	allTags := readTestTags(t, "testdata/tags.txt")

	encoder := NewTagEncoder()

	var tagIndices []int

	for _, tags := range allTags {
		tagIndex := encoder.Encode(tags)
		tagIndices = append(tagIndices, tagIndex)
	}

	b64, err := ioutil.ReadFile("testdata/tags_encoded.txt")
	require.NoError(t, err)

	buf, err := base64.StdEncoding.DecodeString(string(b64))
	require.NoError(t, err)

	for i, tagIndex := range tagIndices {
		assert.Equal(t, allTags[i], getTags(buf, tagIndex))
	}
}

func TestUnsafeIterationV1(t *testing.T) {
	// buff = 2
	buf := make([]byte, 2)
	assert.NotPanics(t, func() {
		unsafeIterateV1(buf, 0, func(i, total int, tag []byte) bool { return true })
	})

	// buff = 1
	assert.NotPanics(t, func() {
		unsafeIterateV1(buf[1:], 0, func(i, total int, tag []byte) bool { return true })
	})

	// indx > buff
	buf = make([]byte, 6)
	assert.NotPanics(t, func() {
		unsafeIterateV1(buf, 10, func(i, total int, tag []byte) bool { return true })
	})

	// footerBuffer < 2
	assert.NotPanics(t, func() {
		unsafeIterateV1(buf, 5, func(i, total int, tag []byte) bool { return true })
	})
}

func FuzzIterateV1(f *testing.F) {
	allTags := readTestTags(f, "testdata/tags.txt")
	t := NewTagEncoder()

	for _, tag := range allTags {
		_ = t.Encode(tag)
	}
	buf := t.Buffer()

	f.Add(buf, 1)
	f.Fuzz(FuzzingIterateV1)
}

func FuzzingIterateV1(t *testing.T, buffer []byte, tagIndex int) {
	assert.NotPanics(t, func() {
		unsafeIterateV1(buffer, tagIndex, func(i, total int, tag []byte) bool { return true })
	})

}

func BenchmarkTagEncode(b *testing.B) {
	allTags := readTestTags(b, "testdata/tags.txt")

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t := NewTagEncoder()

		for _, tags := range allTags {
			_ = t.Encode(tags)
		}
	}
}

func readTestTags(t require.TestingT, filename string) [][]string {
	buf, err := ioutil.ReadFile(filename)
	require.NoError(t, err)

	var allTags [][]string
	for _, line := range strings.Split(string(buf), "\n") {
		tags := strings.Split(line, " ")
		allTags = append(allTags, tags)
	}

	return allTags
}
