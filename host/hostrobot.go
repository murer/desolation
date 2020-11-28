package host

import (
	"fmt"
	"log"
	"os"

	"github.com/murer/desolation/message"
)

var currentRid uint64 = 10

func nRid() uint64 {
	ret := currentRid
	currentRid = currentRid + 1
	return ret
}

func HostCommand(msg *message.Message) *message.Message {
	rid := nRid()
	msg.Rid = rid
	log.Printf("Sent: %s", msg.Basic())
	HostSendMsg(msg)
	ret := CaptureRid(rid)
	log.Printf("Received: %s", ret.Basic())
	return ret
}

func checkConn() {
	rid := nRid()
	msg := message.CreateString(message.OpEcho, rid, "checktext")
	msg = HostCommand(msg)
	if msg.PayloadString() != "checktext" {
		log.Fatalf("Wrong: %v", msg)
	}
}

func Start() {
	msg := CaptureRid(1)
	log.Printf("Init: %v", msg)
	checkConn()

	m := msg.PayloadMap()
	host := m["host"]
	port := m["port"]
	SocketConnect(fmt.Sprintf("%s:%s", host, port))

	for {
		hostDataSend()
		hostDataReceived()
	}
}

func hostDataSend() {
	data := SocketRead()
	if data == nil {
		HostCommand(message.Create(message.OpCloseWrite, 0, []byte{}))
		os.Exit(0)
	}
	if len(data) > 0 {
		msg := HostCommand(message.Create(message.OpWrite, 0, data))
		if msg.Op != message.OpOk {
			log.Panicf("communication error: %v", msg)
		}
	}
}

func hostDataReceived() bool {
	msg := HostCommand(message.Create(message.OpRead, 0, []byte{}))
	if msg == nil {
		return false
	}
	if msg.Op != message.OpOk {
		log.Panicf("communication error: %v", msg)
	}
	SocketWrite(msg.Payload)
	return true
}

func handleResponse(msg *message.Message) {
	if msg.Op == message.OpInit {
		handleResponseInit(msg)
	} else {
		log.Panicf("Unknown: %v", msg)
	}
}

func handleResponseInit(msg *message.Message) {
	m := msg.PayloadMap()
	host := m["host"]
	port := m["port"]
	SocketConnect(fmt.Sprintf("%s:%s", host, port))
}
