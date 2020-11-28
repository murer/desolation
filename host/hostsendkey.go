package host

import (
	"context"
	"os/exec"
	"strings"
	"time"

	"github.com/murer/desolation/message"
	"github.com/murer/desolation/util"
)

// xdotool windowactivate --sync 31457295 key l s Return
// xdotool key a b c Return

func HostSendKeys(keys []string) {
	ctx, cancel := context.WithTimeout(context.Background(), 60000*time.Millisecond)
	defer cancel()
	args := append([]string{"key", "--delay", "0"}, keys...)
	cmd := exec.CommandContext(ctx, "xdotool", args...)
	cmd.Start()
	err := cmd.Wait()
	util.Check(err)
}

func HostSendMsg(msg *message.Message) {
	encoded := msg.Encode()
	array := strings.Split(encoded, "")
	for i := 0; i < len(array); i++ {
		if array[i] == "-" {
			array[i] = "minus"
		} else if array[i] == "_" {
			array[i] = "underscore"
		} else if array[i] == "=" {
			array[i] = "equal"
		}
	}
	HostSendKeys([]string{"ctrl+a"})
	HostSendKeys([]string{"BackSpace"})
	time.Sleep(5 * time.Millisecond)
	HostSendKeys(array)
	HostSendKeys([]string{"Return"})
	time.Sleep(5 * time.Millisecond)
}
