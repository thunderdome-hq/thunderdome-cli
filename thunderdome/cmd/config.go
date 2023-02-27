package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	ConfigCmd = &cobra.Command{
		Use:   "c",
		Short: "Current config",
		Long:  "View or reset the current config",
		RunE:  configAction,
	}
	resetFlag = false
)

func init() {

	ConfigCmd.Flags().BoolVar(&resetFlag, "reset", false, "Reset the config file")

	RootCmd.AddCommand(ConfigCmd)

}

// configAction operates on the config
func configAction(cmd *cobra.Command, args []string) error {
	//Switch the command

	log.Infoln("Viewing config")
	log.Infoln("Config: ", config)
	for _, key := range viper.AllKeys() {
		log.Infof("%s: %v", key, viper.Get(key))
	}
	if resetFlag {
		log.Infoln("Resetting config")
		err := os.Remove(config)
		if err != nil {
			return err
		}
	}

	return nil
}
