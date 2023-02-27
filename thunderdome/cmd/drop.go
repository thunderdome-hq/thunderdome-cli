package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/thunderdome-hq/thunderdome-api/api"
)

var (
	dropOptions   = []string{emailFlag, tokenFlag}
	dropTemplates = []string{"drop.md"}

	dropCmd = &cobra.Command{
		Use:   "drop",
		Short: "Drop your ticket",
		Long:  `Drop the ticket you currently have so that someone else can continue working on it.`,
		RunE:  newAction(dropAction, dropOptions, dropTemplates),
	}
)

func init() {
	RootCmd.AddCommand(dropCmd)
}

// dropAction sends a drop command from the CLI to the Thunderdome lib
func dropAction(cmd *cobra.Command, args []string, client api.ThunderdomeClient, credentials *api.Credentials) (any, error) {
	request := &api.DropTicketRequest{
		Credentials: credentials,
	}

	return client.DropTicket(context.Background(), request)
}
