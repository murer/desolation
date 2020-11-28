package message

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"log"

	"github.com/murer/desolation/util"
)

const OpEcho uint8 = 0
const OpOk uint8 = 1
const OpInit uint8 = 2
const OpWrite uint8 = 3
const OpRead uint8 = 4
const OpCloseWrite uint8 = 5
const OpUnknown uint8 = 6

type Message struct {
	Op      uint8
	Rid     uint64
	Payload []byte
}

func (me *Message) PayloadMap() map[string]string {
	ret := map[string]string{}
	json.Unmarshal(me.Payload, &ret)
	return ret
}

func Create(op uint8, rid uint64, payload []byte) *Message {
	return &Message{
		Op:      op,
		Rid:     rid,
		Payload: payload,
	}
}

func CreateString(op uint8, rid uint64, str string) *Message {
	return Create(op, rid, []byte(str))
}

func CreateMap(op uint8, rid uint64, params map[string]string) *Message {
	str, err := json.Marshal(params)
	util.Check(err)
	return Create(op, rid, str)
}

func CreateUnknown(op uint8) *Message {
	return Create(op, 0, []byte{})
}

func (me *Message) Encode() string {
	var buf bytes.Buffer
	binary.Write(&buf, binary.BigEndian, me.Op)
	binary.Write(&buf, binary.BigEndian, me.Rid)
	size := uint16(len(me.Payload))
	binary.Write(&buf, binary.BigEndian, size)
	buf.Write(me.Payload)
	data := buf.Bytes()
	log.Printf("DEBUG ENCODE: %d %d %x %v", len(data), size, data, me)
	return util.B64Enc(data)
}

func Decode(code string) *Message {
	data := util.B64Dec(code)
	ret := &Message{}
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &ret.Op)
	binary.Read(buf, binary.BigEndian, &ret.Rid)
	var size uint16
	binary.Read(buf, binary.BigEndian, &size)
	ret.Payload = util.ReadFully(buf, int(size))
	log.Printf("DEBUG DECODE: %d, %#v", size, ret)
	return ret
}

// type Message struct {
// 	Name    string
// 	Headers map[string]string
// 	Payload string
// }

// func (m *Message) Get(name string) string {
// 	ret := m.Headers[name]
// 	if ret == "" {
// 		log.Panicf("Message header is required: %s", name)
// 	}
// 	return ret
// }

// func (m *Message) GetInt(name string) int {
// 	ret, err := strconv.Atoi(m.Get(name))
// 	util.Check(err)
// 	return ret
// }

// func (m *Message) PayloadEncode(payload []byte) {
// 	m.Payload = util.B64Enc(payload)
// }

// func (m *Message) PayloadDecode() []byte {
// 	return util.B64Dec(m.Payload)
// }

// func Encode(m *Message) string {
// 	ret, err := json.Marshal(m)
// 	util.Check(err)
// 	return string(ret)
// }

// func Decode(msg string) *Message {
// 	ret := &Message{}
// 	err := json.Unmarshal([]byte(msg), ret)
// 	util.Check(err)
// 	return ret
// }
