package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/spf13/cobra"

	"github.com/murer/desolation/guest"
	"github.com/murer/desolation/host"
	"github.com/murer/desolation/message"
	"github.com/murer/desolation/util"
)

var rootCmd *cobra.Command
var clientCmd *cobra.Command
var pipeCmd *cobra.Command

func Config() {
	rootCmd = &cobra.Command{
		Use: "desolation", Short: "Desolation Proxy",
		Version: fmt.Sprintf("%s-%s:%s", runtime.GOOS, runtime.GOARCH, util.Version),
	}
	rootCmd.PersistentFlags().BoolP("quiet", "q", false, "Quiet")
	cobra.OnInitialize(gconf)

	rootCmd.AddCommand(&cobra.Command{
		Use: "version",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf(rootCmd.Version)
			return nil
		},
	})

	configGuest()

	configHost()

	configCap()

	configSend()

	configCheckConn()

}

func gconf() {
	quiet, err := rootCmd.PersistentFlags().GetBool("quiet")
	util.Check(err)
	if quiet {
		log.SetOutput(ioutil.Discard)
	}
}

func configGuest() {
	rootCmd.AddCommand(&cobra.Command{
		Use:  "guest",
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			guest.TargetHost = args[0]
			guest.TargetPort = args[1]
			guest.Start()
			return nil
		},
	})
}

func makeSleep(cmd *cobra.Command) {
	sleep, err := cmd.PersistentFlags().GetInt64("sleep")
	util.Check(err)
	if sleep > 0 {
		log.Printf("You have %d seconds to put the cursor in the guest text input", sleep)
		time.Sleep(time.Duration(sleep) * time.Second)
	}
}

func configHost() {
	cmd := &cobra.Command{
		Use:  "host",
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			makeSleep(cmd)
			td, err := cmd.PersistentFlags().GetString("td")
			util.Check(err)
			host.SendKeyDelay = td
			host.Start()
			return nil
		},
	}
	cmd.PersistentFlags().Int64("sleep", 5, "Time you need to position your cursor on the guest input text")
	cmd.PersistentFlags().String("td", host.SendKeyDelay, "Send Key Delay in millis")
	rootCmd.AddCommand(cmd)

}

func configCap() {
	cmd := &cobra.Command{
		Use:  "cap",
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			makeSleep(cmd)
			msg := host.CaptureText()
			fmt.Printf("%s", msg)
			return nil
		},
	}
	cmd.PersistentFlags().Int64("sleep", 5, "Time you need to position your cursor on the guest input text")
	rootCmd.AddCommand(cmd)

}

func configSend() {
	cmd := &cobra.Command{
		Use:  "send",
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			makeSleep(cmd)
			content := util.ReadAll(os.Stdin)
			host.HostSendMsg(message.Create(message.OpShow, 0, content))
			return nil
		},
	}
	cmd.PersistentFlags().Int64("sleep", 5, "Time you need to position your cursor on the guest input text")
	rootCmd.AddCommand(cmd)
}

func configCheckConn() {
	cmd := &cobra.Command{
		Use:  "checkconn",
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			makeSleep(cmd)
			host.CheckConn()
			log.Printf("Success")
			return nil
		},
	}
	cmd.PersistentFlags().Int64("sleep", 5, "Time you need to position your cursor on the guest input text")
	rootCmd.AddCommand(cmd)
}

func Execute() {
	err := rootCmd.Execute()
	util.Check(err)
}
