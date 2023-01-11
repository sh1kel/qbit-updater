/*
Copyright © 2023 Alexey Elagin <me@sh1kel.com>

*/

package cmd

import (
    "os"
    "qbit-updater/internals/app"
    "qbit-updater/internals/configuration"

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
    log        *logrus.Logger
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
    //log = logrus.New()
    config := configuration.InitConfig(configFile)
    _ = configureLogger(config)
    app.Process(config)
}
