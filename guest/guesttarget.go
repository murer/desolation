package guest

import (
	"net/http"

	"github.com/murer/desolation/message"
)

var TargetHost = "127.0.0.1"
var TargetPort = "22"

func HandleCommandInit(m *message.Message, w http.ResponseWriter, r *http.Request) *message.Message {
	return &message.Message{
		Name: "ok",
		Headers: map[string]string{
			"host": TargetHost,
			"port": TargetPort,
		},
		Payload: "",
	}
}
