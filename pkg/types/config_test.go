package types

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler(t *testing.T) {
	config := GetReplayedConfig()
	assert.Equal(t, config.BufferSizeInMB, 1000)
}
