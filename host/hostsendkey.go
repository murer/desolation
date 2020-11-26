package host

import (
	"context"
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/murer/desolation/message"
	"github.com/murer/desolation/util"
)

// xdotool windowactivate --sync 31457295 key l s Return
// xdotool key a b c Return

func HostSendKeys(keys []string) {
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()
	args := append([]string{"key"}, keys...)
	log.Printf("aaaa %#v", args)
	cmd := exec.CommandContext(ctx, "xdotool", args...)
	cmd.Start()
	err := cmd.Wait()
	util.Check(err)
}

func HostSendMsg(msg *message.Message) {
	encoded := message.Encode(msg)
	array := strings.Split(encoded, "")
	array = append(array, "Return")
	HostSendKeys(array)

}
