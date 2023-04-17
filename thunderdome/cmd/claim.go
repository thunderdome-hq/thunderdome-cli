package cmd

import (
	"context"

	"github.com/thunderdome-hq/thunderdome-api/api"

	"github.com/spf13/cobra"
)

var (
	claimOptions   = []string{emailFlag, tokenFlag}
	claimTemplates = []string{"claim.md", "ticket.md"}

	claimCmd = &cobra.Command{
		Use:   "claim <ticket id>",
		Short: "Claim a ticket",
		Long:  `Claim a ticket from outsource list by providing its identifier.`,
		Args:  cobra.ExactArgs(1),
		RunE:  newAction(claimAction, claimOptions, claimTemplates),
	}
)

func init() {
	RootCmd.AddCommand(claimCmd)
}

// claimAction sends a claim command from the CLI to the Thunderdome server
func claimAction(cmd *cobra.Command, args []string, client api.ThunderdomeClient, credentials *api.Credentials) (any, error) {
	request := &api.ClaimTicketRequest{
		Request: &api.Request{Credentials: credentials},
		Id:      args[0],
	}

	return client.ClaimTicket(context.Background(), request)
}
