package cmd

import (
	"github.com/adrg/xdg"
	log "github.com/sirupsen/logrus"
	"github.com/thunderdome-hq/thunderdome-cli/thunderdome/render"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	// flags
	hostFlag   = "host"
	portFlag   = "port"
	outputFlag = "output"
	emailFlag  = "email"
	tokenFlag  = "token"
	githubFlag = "github"

	// defaults
	defaultHost   = "thunderdome.hummy.dev"
	defaultPort   = 80
	defaultOutput = "text"
	defaultEmail  = ""
	defaultGithub = ""
	defaultToken  = ""

	// config
	prefix        = "TH"
	defaultConfig = "thunderdome/config.yaml"
)

var (
	config  = "./config.yaml"
	RootCmd = &cobra.Command{
		Use:              "thunderdome [command]",
		Short:            "Thunderdome CLI",
		Long:             "Thunderdome CLI",
		SilenceUsage:     true,
		SilenceErrors:    true,
		PersistentPreRun: preRun,
		RunE:             run,
		PostRunE:         postRun,
	}
)

func init() {
	guide, err := render.Markdown(os.Args[0], "guide.md")
	if err != nil {
		log.Debugln("Unable to render thunderdome guide")
	} else {
		RootCmd.Long = guide
	}

	file, err := xdg.ConfigFile(defaultConfig)
	if err != nil {
		log.Debugln("Unable to set optimal config path, using current directory.")
	} else {
		config = file
	}

	flags := RootCmd.PersistentFlags()

	flags.String(hostFlag, defaultHost, "server host")
	flags.Int(portFlag, defaultPort, "server port")
	flags.String(outputFlag, defaultOutput, "output format")
	flags.String(emailFlag, defaultEmail, "user defaultEmail address")
	flags.String(githubFlag, defaultGithub, "defaultGithub username")
	flags.String(tokenFlag, defaultToken, "user defaultToken")

	env := []string{hostFlag, portFlag, outputFlag, emailFlag, githubFlag, tokenFlag}
	for _, key := range env {
		flag := flags.Lookup(key)
		if viper.BindPFlag(key, flag) != nil {
			log.Debugf("Unable to bild flag %s to %v", key, flag)
		}
	}

	viper.AutomaticEnv()
	viper.SetEnvPrefix(prefix)

	if viper.BindEnv(env...) != nil {
		log.Debugln("Unable to bind environment")
	}
}

func preRun(*cobra.Command, []string) {
	viper.SetConfigFile(config)
	if os.IsNotExist(viper.ReadInConfig()) {
		log.Debugln("Unable to find configuration file, using defaults.")
	}
}

func run(*cobra.Command, []string) error {
	log.Debugln("Environment:")
	for _, key := range viper.AllKeys() {
		log.Debugf("%s: %v", key, viper.Get(key))
	}
	return nil
}

func postRun(*cobra.Command, []string) error {
	return viper.WriteConfigAs(config)
}
