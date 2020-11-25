package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"runtime"

	"github.com/spf13/cobra"

	"github.com/murer/desolation/guest"
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

func Execute() {
	err := rootCmd.Execute()
	util.Check(err)
}
