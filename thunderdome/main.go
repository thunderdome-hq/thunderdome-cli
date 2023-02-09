package main

import (
	"github.com/thunderdome-hq/thunderdome-cli/thunderdome/cmd"
	"google.golang.org/grpc/status"
	"os"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		stat := status.Convert(err)

		cmd.RootCmd.PrintErrf("error: %v: %v\n", stat.Code(), stat.Message())
		cmd.RootCmd.PrintErrln("Please contact Thunderdome admins if you think this is a bug.")

		os.Exit(-1)
	}
}
