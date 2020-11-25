package host

import (
	"context"
	"io"
	"log"
	"os/exec"
	"time"

	"github.com/murer/desolation/message"
	"github.com/murer/desolation/util"
)

func screenshot() []byte {
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()
	out, err := exec.CommandContext(ctx, "import", "-window", "root", "png:-").Output()
	util.Check(err)
	return out
}

func writeQRCOdeStdin(img []byte, stdin io.WriteCloser) {
	defer stdin.Close()
	n, err := stdin.Write(img)
	util.Check(err)
	if n != len(img) {
		log.Panicf("wrong, expected: %d, but was %d", len(img), n)
	}
}

func parseQRCode(img []byte) string {
	ctx, cancel := context.WithTimeout(context.Background(), 20000*time.Millisecond)
	defer cancel()
	cmd := exec.CommandContext(ctx, "zbarimg", "--raw", "png:-")
	stdin, err := cmd.StdinPipe()
	util.Check(err)
	go writeQRCOdeStdin(img, stdin)
	out, err := cmd.Output()
	if err != nil {
		exiterr, ok := err.(*exec.ExitError)
		if ok {
			if exiterr.ExitCode() == 4 {
				return ""
			}
		}
		util.Check(err)
	}
	return string(out)
}

func Capture() *message.Message {
	img := screenshot()
	text := parseQRCode(img)
	if text == "" {
		log.Printf("x")
		return &message.Message{Name: "nocode", Payload: ""}
	}
	log.Printf("y")
	return message.Decode(text)
}

func CaptureRid(rid string) *message.Message {
	retries := util.TimeRetryCreate(10)
	for {
		msg := Capture()
		if msg != nil && msg.Get("rid") == rid {
			return msg
		}
		if retries.Expired() {
			log.Panicf("Timeout waiting for qrcode reply: %s", rid)
		}
		time.Sleep(100 * time.Millisecond)
	}
}
