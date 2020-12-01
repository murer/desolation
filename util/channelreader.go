package util

import (
	"io"
	"log"
	"net"

	"github.com/murer/desolation/util/queue"
)

const DESC_ERR_NONE = 0
const DESC_ERR_OTHER = 1
const DESC_ERR_EOF = 2
const DESC_ERR_TIMEOUT = 3

func ReaderDescError(err error) int {
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
	return DESC_ERR_OTHER
}

func ChannelReader(in io.Reader, bufferLen int) *queue.Queue {
	q := queue.New(1)
	go func() {
		for {
			buf := make([]byte, bufferLen)
			n, err := in.Read(buf)
			derr := ReaderDescError(err)
			if derr == DESC_ERR_EOF {
				log.Print("Reader EOF...")
				q.Put("closed")
				return
			}
			if derr != DESC_ERR_TIMEOUT {
				Check(err)
			}
			q.Put(buf[:n])
		}
	}()
	return q
}
