package util

import (
	"io"
	"log"
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
	// netErr, ok := err.(net.Error)
	// if ok && netErr.Timeout() {
	// 	return DESC_ERR_TIMEOUT
	// }
	// return DESC_ERR_OTHER
	return DESC_ERR_TIMEOUT
}

type ChannelTuple struct {
	Data []byte
	Open bool
}

func ChannelReader(in io.Reader, bufferLen int) chan ChannelTuple {
	ch := make(chan ChannelTuple)
	go func() {
		defer func() {
			log.Print("PPPP")
			close(ch)
		}()
		for {
			buf := make([]byte, bufferLen)
			n, err := in.Read(buf)
			derr := ReaderDescError(err)
			if derr == DESC_ERR_EOF {
				log.Print("Reader EOF...")
				ch <- ChannelTuple{nil, false}
				return
			}
			if derr != DESC_ERR_TIMEOUT {
				Check(err)
			}
			ch <- ChannelTuple{buf[:n], true}
		}
	}()
	return ch
}
