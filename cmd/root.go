/*
Copyright © 2023 Alexey Elagin <me@sh1kel.com>

*/

package cmd

import (
	"github.com/sh1kel/qbit-updater/internals/app"
	"github.com/sh1kel/qbit-updater/internals/configuration"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
	rootCmd = &cobra.Command{
		Use:   "qbit-updater",
		Short: "update qbittorrent tasks",
		Long:  `blabla`,
		Run:   run,
	}
	configFile string
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "",
		"configuration file path (default is .qbit-updater.yaml)")
}

func configureLogger(config *configuration.Config) error {
	config.Logger = logrus.New()
	logLevel, err := logrus.ParseLevel(config.Log.LogLevel)
	if err != nil {
		return err
	}
	config.Logger.SetLevel(logLevel)
	switch config.Log.LogFormat {
	case "json":
		config.Logger.SetFormatter(&logrus.JSONFormatter{})
	default:
		config.Logger.SetFormatter(&logrus.TextFormatter{})
	}
	config.Logger.Infof("log format: %s", config.Log.LogFormat)
	config.Logger.Infof("log level: %s", logLevel)

	return nil
}

func run(cmd *cobra.Command, args []string) {
	config := configuration.InitConfig(configFile)
	err := configureLogger(config)
	if err != nil {
		config.Logger.Fatal(err)
	}
	app.Process(config)
}
