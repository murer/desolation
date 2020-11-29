package guest

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/murer/desolation/message"
	"github.com/murer/desolation/util"
)

var Out io.WriteCloser = os.Stdout
var In io.ReadCloser = os.Stdin

const DESC_ERR_NONE = 0
const DESC_ERR_OTHER = 1
const DESC_ERR_EOF = 2
const DESC_ERR_TIMEOUT = 3

func DescError(err error) int {
	if err == nil {
		return DESC_ERR_NONE
	}
	if err == io.EOF {
		return DESC_ERR_EOF
	}
	// netErr, ok := err.(net.Error)
	// if ok && netErr.Timeout() {
	// 	return DESC_ERR_TIMEOUT
	// }
	// return DESC_ERR_OTHER
	return DESC_ERR_TIMEOUT
}

func HandleCommandWrite(m *message.Message, w http.ResponseWriter, r *http.Request) *message.Message {
	payload := m.Payload
	Out.Write(payload)
	return message.Create(message.OpOk, 0, []byte{})
}

func HandleCommandRead(m *message.Message, w http.ResponseWriter, r *http.Request) *message.Message {
	buf := make([]byte, 256)
	os.Stdin.SetReadDeadline(time.Now().Add(1 * time.Millisecond))
	n, err := os.Stdin.Read(buf)
	derr := DescError(err)
	if derr == DESC_ERR_EOF {
		log.Print("Stdin EOF...")
		return nil
	}
	if derr != DESC_ERR_TIMEOUT {
		util.Check(err)
	}
	buf = buf[:n]
	return message.Create(message.OpOk, 0, buf)
}

func HandleCommandCW(m *message.Message, w http.ResponseWriter, r *http.Request) *message.Message {
	Out.Close()
	return message.Create(message.OpOk, 0, []byte{})
}
