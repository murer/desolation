package host

import (
	"fmt"
	"log"
	"os"

	"github.com/murer/desolation/message"
	"github.com/murer/desolation/util"
)

var currentRid uint32 = 10

func nRid() uint32 {
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

func CheckConn() *message.Message {
	data := util.Rand(256)
	rid := nRid()
	msg := message.Create(message.OpEcho, rid, data)
	msg = HostCommand(msg)
	if len(msg.Payload) != len(data) {
		log.Fatalf("Wrong: %v", msg)
	}
	for index, element := range data {
		if element != msg.Payload[index] {
			log.Fatalf("Wrong, idx: %d, exp: %d, but was: %d", index, element, msg.Payload[index])
		}
	}
	return HostCommand(message.Create(message.OpInit, 0, []byte{}))
}

func Start() {
	msg := CheckConn()
	log.Printf("Init: %v", msg)

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

func hostDataReceived() {
	msg := HostCommand(message.Create(message.OpRead, 0, []byte{}))
	if msg == nil {
		return
	}
	if msg.Op == message.OpReaderClosed {
		SocketReaderClose()
		return
	}
	if msg.Op != message.OpOk {
		log.Panicf("communication error: %v", msg)
	}
	SocketWrite(msg.Payload)
	return
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
