package message_test

import (
	"testing"

	"github.com/murer/desolation/message"
	"github.com/stretchr/testify/assert"
)

func TestMessage(t *testing.T) {
	original := message.CreateMap(message.OpRead, 1, map[string]string{"x": "murer"})
	assert.Equal(t, map[string]string{"x": "murer"}, original.PayloadMap())
	code := original.Encode()
	assert.Equal(t, "AwAAAAAAAAABAA17IngiOiJtdXJlciJ9", code)
	msg := message.Decode(code)
	assert.Equal(t, original, msg)
}
