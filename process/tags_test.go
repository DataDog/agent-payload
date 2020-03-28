package process

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"math"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TagSerdeTestSuite struct {
	suite.Suite
	encoder TagEncoder
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

	assert.Equal(suite.T(), []string{"one", "two", "three"}, getTags(buf, a))
	assert.Empty(suite.T(), getTags(buf, b))
	assert.Equal(suite.T(), []string{"four", "five"}, getTags(buf, c))
	assert.Equal(suite.T(), []string{"six"}, getTags(buf, d))
	assert.Equal(suite.T(), []string{"seven", "eight", "nine", "ten"}, getTags(buf, e))
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
		assert.Equal(suite.T(), allTags[i], getTags(encoder.Buffer(), tagIndex))

		var iterated []string
		iterateTags(encoder.Buffer(), tagIndex, func(tag string) {
			iterated = append(iterated, tag)
		})
		assert.Equal(suite.T(), allTags[i], iterated)

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
