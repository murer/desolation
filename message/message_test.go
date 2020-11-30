package message_test

import (
	"log"
	"testing"

	"github.com/murer/desolation/message"
	"github.com/stretchr/testify/assert"
)

func TestMessage(t *testing.T) {
	// original := message.CreateMap(message.OpRead, 1, map[string]string{"x": "murer"})
	original := message.Decode("AQAAABYAWGExAAAABG5vbmUAAAAEbm9uZQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAsBh4AAAAg2fdmBPDfKDrh4aDGRcgvUctBZHEsc-GBBJn0xPcc-BMAAAAAAAAH8yhK")
	log.Printf("msg: %s", original.Basic())
	assert.Equal(t, map[string]string{"x": "murer"}, original.Payload)
	code := original.Encode()
	assert.Equal(t, "BAAAAAEADXsieCI6Im11cmVyIn0ej6o0", code)
	msg := message.Decode(code)
	assert.Equal(t, original, msg)
}
