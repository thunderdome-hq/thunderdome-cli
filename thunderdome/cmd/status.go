package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/thunderdome-hq/thunderdome-api/api"
)

var (
	statusOptions   = []string{emailFlag, tokenFlag}
	statusTemplates = []string{"status.md", "user.md", "ticket.md"}
)

func init() {
	RootCmd.AddCommand(statusCmd)
}

// statusAction sends a status command from the CLI to the Thunderdome server
func statusAction(cmd *cobra.Command, args []string, client api.ThunderdomeClient, credentials *api.Credentials) (any, error) {
	request := &api.GetStatusRequest{
		Credentials: credentials,
	}
	return client.GetStatus(context.Background(), request)
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Get status",
	Long:  fmt.Sprintf(`Get status of current user and claimed ticket. Required options: %s`, strings.Join(statusOptions, ", ")),
	RunE:  newAction(statusAction, statusOptions, statusTemplates),
}
