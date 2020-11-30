package message_test

import (
	"log"
	"testing"

	"github.com/murer/desolation/message"
	"github.com/stretchr/testify/assert"
)

func TestMessage(t *testing.T) {
	original := message.CreateMap(message.OpRead, 1, map[string]string{"x": "murer"})
	//original := message.Decode("aAAAAAsACWNoZWNrdGV4dKQf4s8")
	log.Printf("msg: %s", original.Basic())
	assert.Equal(t, map[string]string{"x": "murer"}, original.PayloadMap())
	code := original.Encode()
	assert.Equal(t, "BAAAAAEADXsieCI6Im11cmVyIn0ej6o0", code)
	msg := message.Decode(code)
	assert.Equal(t, original, msg)
}
