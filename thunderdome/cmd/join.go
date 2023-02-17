package cmd

import (
	"context"

	"github.com/thunderdome-hq/thunderdome-api/api"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	joinOptions   = []string{emailFlag, githubFlag}
	joinTemplates = []string{"join.md", "user.md"}

	joinCmd = &cobra.Command{
		Use:   "join",
		Short: "Join Thunderdome",
		Long:  `Enroll in Thunderdome by providing your defaultEmail and defaultGithub username. Administrators will be notified and asked to accept or reject you, after which you will receive an defaultEmail. The first time you run this command, you can specify your credentials with the --defaultEmail and --defaultGithub flags. After that, they will be stored in your defaultConfig file and can be omitted.`,
		RunE:  newAction(joinAction, joinOptions, joinTemplates),
	}
)

func init() {
	RootCmd.AddCommand(joinCmd)
}

// joinAction sends a join command from the CLI to the Thunderdome server
func joinAction(cmd *cobra.Command, args []string, client api.ThunderdomeClient, credentials *api.Credentials) (any, error) {
	request := &api.JoinUserRequest{
		Email:  viper.GetString(emailFlag),
		Github: viper.GetString(githubFlag),
	}

	response, err := client.JoinUser(context.Background(), request)
	if err != nil {
		return response, err
	}

	viper.Set(tokenFlag, response.Token)

	err = viper.WriteConfig()
	if err != nil {
		return response, err
	}

	cmd.PrintErrln("Storing credentials in", viper.ConfigFileUsed())

	return response, nil
}
