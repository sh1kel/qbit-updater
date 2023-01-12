package app

import (
	"github.com/sh1kel/qbit-updater/internals/configuration"
	"github.com/sh1kel/qbit-updater/internals/tclient"
)

func Process(config *configuration.Config) {
	log := config.Logger
	qc := tclient.New("http://192.168.1.2:8085", "1", "2", log)
	err := qc.Connect()
	if err != nil {
		log.Error(err)
	}
	version, _ := qc.GetVersion()
	log.Infof("qB version: %s", version)
	torrents, err := qc.GetAllTorrents(nil)
	if err != nil {
		log.Error(err)
	}
	for _, t := range torrents {
		ti, _ := qc.GetTorrentInfo(t.Hash)
		log.Debug(*ti)

	}
	err = qc.DownloadFromFile("111.torrent", map[string]string{})
	if err != nil {
		log.Error(err)
	}

}
