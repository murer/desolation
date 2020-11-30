package util_test

import (
	"testing"

	"github.com/murer/desolation/util"
	"github.com/stretchr/testify/assert"
)

func TestCrc32(t *testing.T) {
	assert.Equal(t, uint32(0xd87f7e0c), util.Crc32([]byte("test")))
	assert.Equal(t, uint32(0xd9583520), util.Crc32([]byte("other")))
}
