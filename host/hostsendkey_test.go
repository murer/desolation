package host_test

import (
	"testing"

	"github.com/murer/desolation/host"
	"github.com/murer/desolation/message"
)

func TestSendkey(t *testing.T) {
	//host.HostSendKeys([]string{"{"})
}

func TestSendMsg(t *testing.T) {
	host.HostSendMsg(&message.Message{
		Name:    "echo",
		Headers: map[string]string{"rid": "x"},
		Payload: "y",
	})
}
