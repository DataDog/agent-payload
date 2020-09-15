package process

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseIPAddress(t *testing.T) {
	ip := ParseIPAddress("1.1.1.1")

	assert.Equal(t, "1.1.1.1", ip.String())
}
