package app

import (
    "qbit-updater/internals/configuration"
    "qbit-updater/internals/tclient"
)

func Process(config *configuration.Config) {
    log := config.Logger
    qc := tclient.New("http://192.168.1.2:8082", "1", "2", log)
    err := qc.Connect()
    if err != nil {
        log.Error(err)
    }

}
