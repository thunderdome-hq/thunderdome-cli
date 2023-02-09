package cmd

import (
	"fmt"
	"github.com/thunderdome-hq/thunderdome-cli/thunderdome/render"
	"os"
	"strings"

	"github.com/adrg/xdg"
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tcnksm/go-gitconfig"
)

var (
	cfgFile string
	gitCfg  bool

	output        string
	outputFormats = []string{"text", "json"}
)

const (
	hostKey   = "host"
	portKey   = "port"
	emailKey  = "email"
	tokenKey  = "token"
	githubKey = "github"
)

const (
	defaultHost = "thunderdome-api.hummy.dev"
	defaultPort = 8080
)

var RootCmd = &cobra.Command{
	Use:           "thunderdome [command]",
	Short:         "Thunderdome CLI",
	Long:          "Thunderdome CLI",
	SilenceUsage:  true,
	SilenceErrors: true,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// Read config
		viper.SetConfigFile(cfgFile)
		err := viper.ReadInConfig()
		if os.IsNotExist(err) {
			log.Debugln("Unable to find configuration file, using defaults.")
		}

		// Check if options should be read from gitconfig
		if gitCfg {
			log.Debugln("Reading options from gitconfig")

			email, err := gitconfig.Email()
			if err != nil {
				return err
			}
			if email != "" {
				viper.Set(emailKey, email)
			}

			github, _ := gitconfig.GithubUser()
			if err != nil {
				return nil
			}
			if github != "" {
				viper.Set(githubKey, github)
			}
		}

		// Ensure that output format is valid
		validFormat := false
		for _, format := range outputFormats {
			if output == format {
				validFormat = true
				break
			}
		}

		if !validFormat {
			return fmt.Errorf("invalid output format: %s", output)
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	// Set long description
	guide, err := render.Markdown(os.Args[0], "guide.md")
	if err != nil {
		fmt.Println(err)
		log.Debugln("Could not render guide")
	} else {
		RootCmd.Long = guide
	}

	defaultCfg, err := xdg.ConfigFile("thunderdome/config.yaml")
	if err != nil {
		fmt.Println(err)
		log.Debugln("Could not determine config file path, looking in current directory.")
		defaultCfg = "./thunderdome.yaml"
	}

	// Set defaults
	viper.SetDefault(hostKey, defaultHost)
	viper.SetDefault(portKey, defaultPort)

	// Config file
	RootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", defaultCfg, "config file")

	// Config flags
	RootCmd.PersistentFlags().String("host", "", "server host")
	RootCmd.PersistentFlags().Int("port", 0, "server port")
	RootCmd.PersistentFlags().String("email", "", "user email address")
	RootCmd.PersistentFlags().String("github", "", "user github account")
	RootCmd.PersistentFlags().String("token", "", "user access token")

	RootCmd.PersistentFlags().BoolVar(&gitCfg, "gitconfig", false, "read options from gitconfig")
	RootCmd.PersistentFlags().StringVarP(&output, "output", "o", "text", fmt.Sprintf("Output format {%s}", strings.Join(outputFormats, ", ")))

	// Bind flags
	viper.BindPFlag(hostKey, RootCmd.PersistentFlags().Lookup("host"))
	viper.BindPFlag(portKey, RootCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag(emailKey, RootCmd.PersistentFlags().Lookup("email"))
	viper.BindPFlag(tokenKey, RootCmd.PersistentFlags().Lookup("token"))
	viper.BindPFlag(githubKey, RootCmd.PersistentFlags().Lookup("github"))

	// Bind env vars
	viper.AutomaticEnv()
	viper.SetEnvPrefix("TD")
	viper.BindEnv(hostKey, portKey, emailKey, tokenKey, githubKey)
}
