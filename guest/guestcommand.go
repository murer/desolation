package guest

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/murer/desolation/message"
	"github.com/murer/desolation/util"
	"github.com/murer/desolation/util/queue"
)

var Out io.WriteCloser
var In *queue.Queue

func GuestInOutInit() {
	if In != nil {
		log.Print("In and out are alreadt setted")
		return
	}
	In = util.ChannelReader(os.Stdin, 256)
	Out = os.Stdout
}

func HandleCommandWrite(m *message.Message, w http.ResponseWriter, r *http.Request) *message.Message {
	payload := m.Payload
	Out.Write(payload)
	return message.Create(message.OpOk, 0, []byte{})
}

func HandleCommandRead(m *message.Message, w http.ResponseWriter, r *http.Request) *message.Message {
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
