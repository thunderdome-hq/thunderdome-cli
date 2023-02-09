package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/thunderdome-hq/thunderdome-api/api"
)

var (
	unclaimed bool
)

var (
	listOptions   = []string{emailKey, tokenKey}
	listTemplates = []string{"list.md", "ticket.md"}
)

func init() {
	listCmd.Flags().BoolVar(&unclaimed, "unclaimed", false, "Show only unclaimed tickets")

	RootCmd.AddCommand(listCmd)
}

// listAction sends a list command from the CLI to the Thunderdome server
func listAction(cmd *cobra.Command, _ []string, client api.ThunderdomeClient, credentials *api.Credentials) (any, error) {
	request := &api.ListTicketsRequest{
		Credentials: credentials,
	}

	response, err := client.ListTickets(context.Background(), request)
	if err != nil {
		return response, err
	}

	var tickets []*api.Ticket
	for _, ticket := range response.Tickets {
		if unclaimed && ticket.User != "" {
			continue
		}

		tickets = append(tickets, ticket)
	}
	response.Tickets = tickets

	return response, nil
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List tickets",
	Long:  `List tickets in Thunderdome using access token.`,
	RunE:  newAction(listAction, listOptions, listTemplates),
}
