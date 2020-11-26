package host

import (
	"fmt"
	"log"
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

func hostDataReceived() {
	msg := HostCommand(&message.Message{
		Name:    "read",
		Headers: map[string]string{},
		Payload: "",
	})
	if msg.Name != "ok" {
		log.Panicf("communication error: %v", msg)
	}
	data := util.B64Dec(msg.Payload)
	SocketWrite(data)
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
