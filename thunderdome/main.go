package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/thunderdome-hq/thunderdome-cli/thunderdome/cmd"
	"google.golang.org/grpc/status"
	"os"
)

const level = log.InfoLevel

func main() {
	log.SetLevel(level)
	if err := cmd.RootCmd.Execute(); err != nil {
		stat := status.Convert(err)
		cmd.RootCmd.PrintErrf(
			"%v error: %v\nPlease contact Thunderdome admins if you think this is a bug.",
			stat.Code(),
			stat.Message())

		os.Exit(-1)
	}
}
