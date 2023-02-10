package app

import (
	"github.com/sh1kel/qbit-updater/internals/configuration"
	"github.com/sh1kel/qbit-updater/internals/tclient"
)

func SetConfigOptions(options map[string]string, config *configuration.Config) {
	log := config.Logger
	for _, url := range config.Clients.Urls {
		qc := tclient.New(url, "", "", log)
		err := qc.Connect()
		if err != nil {
			log.Error(err)
		}
		version, _ := qc.GetVersion()
		if err != nil {
			log.Error(err)
		}
		log.Infof("[%s] qB version: %s", url, version)
		for k, v := range options {
			log.Infof("set %s: %s", k, v)
		}
		err = qc.SetApplicationPreferences(options)
		if err != nil {
			log.Error(err)
		}
	}

}
