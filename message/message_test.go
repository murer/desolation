package message_test

import (
	"testing"

	"github.com/murer/desolation/message"
	"github.com/stretchr/testify/assert"
)

func TestMessage(t *testing.T) {
	original := &message.Message{
		Name:    "nm",
		Headers: map[string]string{"foo": "1", "bar": "2"},
		Payload: "bXVyZXI=",
	}
	assert.Equal(t, "1", original.Get("foo"))
	assert.Equal(t, 2, original.GetInt("bar"))
	assert.Equal(t, []byte("murer"), original.PayloadDecode())
	original.PayloadEncode([]byte{1, 2})
	assert.Equal(t, []byte{1, 2}, original.PayloadDecode())

	// secret := []byte("12345678901234561234567890123456")
	// buf := MessageEnc(secret, original)
	// msg := MessageDec(secret, buf)
	// assert.Equal(t, original, msg)
	// assert.Equal(t, 48, len(buf))
}

func TestParse(t *testing.T) {
	original := &message.Message{
		Name:    "nm",
		Headers: map[string]string{"foo": "1", "bar": "2"},
		Payload: "bXVyZXI=",
	}

	encoded := message.Encode(original)
	assert.Equal(t, "{\"Name\":\"nm\",\"Headers\":{\"bar\":\"2\",\"foo\":\"1\"},\"Payload\":\"bXVyZXI=\"}", encoded)

	msg := message.Decode(encoded)
	assert.Equal(t, original, msg)
}
