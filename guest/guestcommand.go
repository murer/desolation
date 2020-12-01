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

func HandleCommandWrite(m *message.Message, w http.ResponseWriter, r *http.Request) *message.Message {
	payload := m.Payload
	Out.Write(payload)
	return message.Create(message.OpOk, 0, []byte{})
}

func HandleCommandRead(m *message.Message, w http.ResponseWriter, r *http.Request) *message.Message {
	log.Print("xxx2")
	content := In.Shift()
	if content == nil {
		return message.Create(message.OpOk, 0, []byte{})
	}
	block, ok := content.([]byte)
	if ok {
		return message.Create(message.OpOk, 0, block)
	}
	closed, ok := content.(string)
	if !ok || closed != "closed" {
		log.Panicf("Wrong: %#v", content)
	}
	log.Print("Stdin EOF...")
	return nil
}

func HandleCommandCW(m *message.Message, w http.ResponseWriter, r *http.Request) *message.Message {
	Out.Close()
	return message.Create(message.OpOk, 0, []byte{})
}
