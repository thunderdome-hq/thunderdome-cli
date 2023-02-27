package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/thunderdome-hq/thunderdome-api/api"
)

var (
	leaveOptions   = []string{emailFlag, tokenFlag}
	leaveTemplates = []string{"leave.md"}

	leaveCmd = &cobra.Command{
		Use:   "leave",
		Short: "Leave Thunderdome",
		Long:  `Leave Thunderdome by providing defaultEmail and defaultToken. This will revoke your access and you will have to rejoin again.`,
		RunE:  newAction(leaveAction, leaveOptions, leaveTemplates),
	}
)

func init() {
	RootCmd.AddCommand(leaveCmd)
}

// leaveAction sends a leave command from the CLI to the Thunderdome lib
func leaveAction(cmd *cobra.Command, args []string, client api.ThunderdomeClient, credentials *api.Credentials) (any, error) {
	request := &api.LeaveUserRequest{
		Credentials: credentials,
	}

	return client.LeaveUser(context.Background(), request)
}
