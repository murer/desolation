package main

import (
	"syscall"

	"github.com/murer/desolation/cmd"
	"github.com/murer/desolation/util"
)

func main() {
	err := syscall.SetNonblock(0, true)
	util.Check(err)
	cmd.Config()
	cmd.Execute()

	// guest.Start()
}
