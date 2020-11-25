package guest

import (
	"io"
	"net/http"
	"os"

	"github.com/murer/desolation/message"
	"github.com/murer/desolation/util"
)

var Out io.WriteCloser = os.Stdout
var In io.ReadCloser = os.Stdin

func HandleCommandWrite(m *message.Message, w http.ResponseWriter, r *http.Request) *message.Message {
	payload := m.PayloadDecode()
	Out.Write(payload)
	return &message.Message{Name: "ok", Headers: map[string]string{}, Payload: ""}
}

func HandleCommandRead(m *message.Message, w http.ResponseWriter, r *http.Request) *message.Message {
	buf := make([]byte, 4)
	n, err := In.Read(buf)
	util.Check(err)
	buf = buf[:n]
	return &message.Message{Name: "ok", Headers: map[string]string{}, Payload: util.B64Enc(buf)}
}

func HandleCommandCW(m *message.Message, w http.ResponseWriter, r *http.Request) *message.Message {
	Out.Close()
	return &message.Message{Name: "ok", Headers: map[string]string{}, Payload: ""}
}
