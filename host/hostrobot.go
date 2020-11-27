package host

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/murer/desolation/message"
	"github.com/murer/desolation/util"
)

var currentRid int64 = 1

func nRid() string {
	ret := currentRid
	currentRid = currentRid + 1
	return strconv.FormatInt(ret, 10)
}

func HostCommand(msg *message.Message) *message.Message {
	rid := nRid()
	msg.Headers["rid"] = rid
	log.Printf("Sent: %v", msg)
	HostSendMsg(msg)
	ret := CaptureRid(rid)
	log.Printf("Received: %v", ret)
	return ret
}

func checkConn() {
	rid := nRid()
	msg := &message.Message{
		Name:    "echo",
		Headers: map[string]string{"rid": rid},
		Payload: "checktext",
	}
	msg = HostCommand(msg)
	if msg.Payload != "checktext" {
		log.Fatalf("Wrong: %v", msg)
	}
}

func Start() {
	msg := CaptureRid("init")
	log.Printf("Init: %v", msg)
	checkConn()

	host := msg.Get("host")
	port := msg.Get("port")
	SocketConnect(fmt.Sprintf("%s:%s", host, port))

	for {
		hostDataSend()
		hostDataReceived()
	}
}

func hostDataSend() {
	data := SocketRead()
	if data == nil {
		HostCommand(&message.Message{
			Name:    "cw",
			Headers: map[string]string{},
			Payload: "",
		})
		os.Exit(0)
	}
	if len(data) > 0 {
		msg := HostCommand(&message.Message{
			Name:    "write",
			Headers: map[string]string{},
			Payload: util.B64Enc(data),
		})
		if msg.Name != "ok" {
			log.Panicf("communication error: %v", msg)
		}
	}
}

func hostDataReceived() bool {
	msg := HostCommand(&message.Message{
		Name:    "read",
		Headers: map[string]string{},
		Payload: "",
	})
	if msg == nil {
		return false
	}
	if msg.Name != "ok" {
		log.Panicf("communication error: %v", msg)
	}
	data := util.B64Dec(msg.Payload)
	SocketWrite(data)
	return true
}

func handleResponse(msg *message.Message) {
	if msg.Name == "init" {
		handleResponseInit(msg)
	} else {
		log.Panicf("Unknown: %v", msg)
	}
}

func handleResponseInit(msg *message.Message) {
	host := msg.Get("host")
	port := msg.Get("port")
	SocketConnect(fmt.Sprintf("%s:%s", host, port))
}
