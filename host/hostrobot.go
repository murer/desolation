package host

import (
	"log"
	"strconv"

	"github.com/murer/desolation/message"
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
}
