package cmd

import (
	"context"

	"github.com/thunderdome-hq/thunderdome-api/api"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	joinOptions   = []string{emailKey, githubKey}
	joinTemplates = []string{"join.md", "user.md"}
)

func init() {
	RootCmd.AddCommand(joinCmd)
}

// joinAction sends a join command from the CLI to the Thunderdome server
func joinAction(cmd *cobra.Command, args []string, client api.ThunderdomeClient, credentials *api.Credentials) (any, error) {
	request := &api.JoinUserRequest{
		Email:  viper.GetString(emailKey),
		Github: viper.GetString(githubKey),
	}

	response, err := client.JoinUser(context.Background(), request)
	if err != nil {
		return response, err
	}

	viper.Set(tokenKey, response.Token)

	err = viper.WriteConfig()
	if err != nil {
		return response, err
	}

	cmd.PrintErrln("Storing credentials in", viper.ConfigFileUsed())

	return response, nil
}

var joinCmd = &cobra.Command{
	Use:   "join",
	Short: "Join Thunderdome",
	Long:  `Enroll in Thunderdome by providing your email and github username. Administrators will be notified and asked to accept or reject you, after which you will receive an email. The first time you run this command, you can specify your credentials with the --email and --github flags. After that, they will be stored in your config file and can be omitted.`,
	RunE:  newAction(joinAction, joinOptions, joinTemplates),
}
