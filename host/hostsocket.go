package host

import (
	"io"
	"log"
	"net"
	"strings"
	"time"

	"github.com/murer/desolation/util"
)

var conn net.Conn

const DESC_ERR_NONE = 0
const DESC_ERR_OTHER = 1
const DESC_ERR_EOF = 2
const DESC_ERR_TIMEOUT = 3
const DESC_ERR_CLOSED = 4

func DescError(err error) int {
	if err == nil {
		return DESC_ERR_NONE
	}
	if err == io.EOF {
		return DESC_ERR_EOF
	}
	netErr, ok := err.(net.Error)
	if ok && netErr.Timeout() {
		return DESC_ERR_TIMEOUT
	}
	if strings.Contains(err.Error(), "use of closed network connection") {
		return DESC_ERR_CLOSED
	}
	return DESC_ERR_OTHER
}

func SocketConnect(addr string) {
	log.Printf("Connecting %s", addr)
	c, err := net.Dial("tcp", addr)
	util.Check(err)
	log.Printf("connected %s", addr)
	conn = c
}

func SocketRead() []byte {
	conn.SetReadDeadline(time.Now().Add(1 * time.Millisecond))
	buf := make([]byte, 512)
	n, err := conn.Read(buf)
	derr := DescError(err)
	if derr == DESC_ERR_EOF {
		log.Print("Socket EOF...")
		return nil
	}
	if derr != DESC_ERR_TIMEOUT {
		util.Check(err)
	}
	buf = buf[:n]
	return buf
}

func SocketWrite(data []byte) {
	if len(data) == 0 {
		return
	}
	n, err := conn.Write(data)
	util.Check(err)
	if n != len(data) {
		log.Panicf("Wrong len, should be: %d, was: %d", n, len(data))
	}
}

func SocketReaderClose() {
	log.Print("Closing connection")
	conn.Close()
}
