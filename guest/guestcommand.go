package guest

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/murer/desolation/message"
	"github.com/murer/desolation/util"
)

var Out io.WriteCloser = os.Stdout
var In = util.ChannelReader(os.Stdin, 256)

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
	log.Print("xxx2")
	tuple := <-In
	log.Print("xxx3")
	if !tuple.Open {
		log.Print("Stdin EOF...")
		return nil
	}
	return message.Create(message.OpOk, 0, tuple.Data)
}

func HandleCommandCW(m *message.Message, w http.ResponseWriter, r *http.Request) *message.Message {
	Out.Close()
	return message.Create(message.OpOk, 0, []byte{})
}
