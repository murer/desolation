package poc_screenshot

// import -window root png:- > test.png
// xwininfo -tree -root

import (
	"context"
	"os/exec"
	"time"
)

func Screenshot() []byte {
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()
	out, err := exec.CommandContext(ctx, "import", "-window", "root", "png:-").Output()
	if err != nil {
		panic(err)
	}
	return out
}
