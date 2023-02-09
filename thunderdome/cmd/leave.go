package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/thunderdome-hq/thunderdome-api/api"
)

var (
	leaveOptions   = []string{emailKey, tokenKey}
	leaveTemplates = []string{"leave.md"}
)

func init() {
	RootCmd.AddCommand(leaveCmd)
}

// leaveAction sends a leave command from the CLI to the Thunderdome server
func leaveAction(cmd *cobra.Command, args []string, client api.ThunderdomeClient, credentials *api.Credentials) (any, error) {
	request := &api.LeaveUserRequest{
		Credentials: credentials,
	}

	return client.LeaveUser(context.Background(), request)
}

var leaveCmd = &cobra.Command{
	Use:   "leave",
	Short: "Leave Thunderdome",
	Long:  `Leave Thunderdome by providing email and token. This will revoke your access and you will have to rejoin again.`,
	RunE:  newAction(leaveAction, leaveOptions, leaveTemplates),
}
