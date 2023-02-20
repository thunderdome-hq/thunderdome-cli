package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thunderdome-hq/thunderdome-api/api"
	"github.com/thunderdome-hq/thunderdome-cli/thunderdome/render"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

type Action func(cmd *cobra.Command, args []string, client api.ThunderdomeClient, credentials *api.Credentials) (any, error)

func newAction(action Action, options []string, templates []string) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		// Host and defaultPort are required for all commands
		options = append(options, hostFlag, portFlag)

		// Check required arguments
		for _, arg := range options {
			if !viper.IsSet(arg) {
				return api.Error(codes.FailedPrecondition, api.CLIError, "missing argument for %s", arg)
			}
		}

		// Connect to server
		target := fmt.Sprintf("%s:%d", viper.GetString(hostFlag), viper.GetInt(portFlag))

		conn, err := grpc.DialContext(
			context.Background(),
			target,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)

		if err != nil {
			log.Debugln("Could not connect to server:", err)
			return api.Error(codes.Unavailable, api.CLIError, "unable to connect to server target %s", target)
		}

		defer conn.Close()

		// Create client
		client := api.NewThunderdomeClient(conn)

		// Create credentials
		credentials := &api.Credentials{
			Email: viper.GetString(emailFlag),
			Token: viper.GetString(tokenFlag),
		}

		apiar := NewSpinner()
		apiar.Start()

		// Execute action
		res, err := action(cmd, args, client, credentials)

		apiar.Stop()

		if err != nil {
			log.Debugln("Could not execute action:", err)
			return err
		}

		switch viper.GetString(outputFlag) {
		case "json":
			out, err := json.MarshalIndent(res, "", "  ")
			if err != nil {
				log.Debugln("Unable to marshal response:", err)
				return api.Error(codes.Internal, api.CLIError, "unable to display JSON response")
			}

			cmd.Println(out)
		case "text":
			out, err := render.Markdown(res, templates...)
			if err != nil {
				return err
			}

			cmd.Printf(out)
		default:
			cmd.Printf("Unknown defaultOutput format, printing raw response:\n%v", res)
		}

		return nil
	}
}