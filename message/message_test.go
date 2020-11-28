package message_test

import (
	"log"
	"testing"

	"github.com/murer/desolation/message"
	"github.com/stretchr/testify/assert"
)

func TestMessage(t *testing.T) {
	original := message.CreateMap(message.OpRead, 1, map[string]string{"x": "murer"})
	//original := message.Decode("AQAAAAAAAAARAXljb20sdW1hYy02NEBvcGVuc3NoLmNvbSx1bWFjLTEyOEBvcGVuc3NoLmNvbSxobWFjLXNoYTItMjU2LGhtYWMtc2hhMi01MTIsaG1hYy1zaGExAAAA1XVtYWMtNjQtZXRtQG9wZW5zc2guY29tLHVtYWMtMTI4LWV0bUBvcGVuc3NoLmNvbSxobWFjLXNoYTItMjU2LWV0bUBvcGVuc3NoLmNvbSxobWFjLXNoYTItNTEyLWV0bUBvcGVuc3NoLmNvbSxobWFjLXNoYTEtZXRtQG9wZW5zc2guY29tLHVtYWMtNjRAb3BlbnNzaC5jb20sdW1hYy0xMjhAb3BlbnNzaC5jb20saG1hYy1zaGEyLTI1NixobWFjLXNoYTItNTEyLGhtYWMtc2hhMQAAABpub25lLHpsaWJAb3BlbnNzaC5jb20semxpYgAAABpub25lLHpsaWJAb3BlbnNzaC5jb20semxpYgAAAAAAAAAAAAAAAAAAAAAAAA")
	log.Printf("msg: %s", original.Basic())
	assert.Equal(t, map[string]string{"x": "murer2"}, original.PayloadMap())
	code := original.Encode()
	assert.Equal(t, "BAAAAAEADXsieCI6Im11cmVyIn0", code)
	msg := message.Decode(code)
	assert.Equal(t, original, msg)
}
