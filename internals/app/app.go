package app

import (
	"github.com/sh1kel/qbit-updater/internals/configuration"
	"github.com/sh1kel/qbit-updater/internals/forum"
	"github.com/sh1kel/qbit-updater/internals/tclient"
	"os"
)

func Process(config *configuration.Config) {
	log := config.Logger
	fc := forum.New(config)
	err := fc.Init()
	if err != nil {
		log.Fatal(err)
	}
	err = fc.Auth()
	if err != nil {
		log.Fatal(err)
	}
	for _, url := range config.Clients.Urls {
		updateCounter := 0
		qc := tclient.New(url, "", "", log)
		err = qc.Connect()
		if err != nil {
			log.Error(err)
			continue
		}
		version, _ := qc.GetVersion()
		if err != nil {
			log.Error(err)
		}
		log.Infof("[%s] qB version: %s", url, version)

		torrents, err := qc.GetAllTorrents(nil)
		torrentsBeforeClean := len(torrents)
		if err != nil {
			log.Error(err)
			continue
		}
		log.Infof("Torrents count: %d", torrentsBeforeClean)

		for _, t := range torrents {
			tt, err := qc.GetTrackers(t.Hash)
			if err != nil {
				log.Error(err)
				continue
			}
			for _, tracker := range tt {
				if tracker.Status == tclient.TrackerHasBeenContactedButItIsNotWorking &&
					tracker.Msg == tclient.TorrentNotRegistered {
					ti, err := qc.GetTorrentInfo(t.Hash)
					if err != nil {
						log.Error(err)
						continue
					}
					log.Infof("%s > %s: %s", t.Category, t.Name, ti.Comment)
					shortHash, err := qc.GetShortHashFromComment(t.Hash)
					if err != nil {
						log.Error(err)
						continue
					}
					err = fc.GetTorrentFile(shortHash)
					if err != nil {
						log.Error(err)
						continue
					}
					err = qc.DownloadFromFile(fc.GetLastDownloadedFileName(),
						map[string]string{"savepath": t.SavePath, "category": t.Category},
					)
					if err != nil {
						log.Error(err)
						continue
					}
					// Updating counter
					updateCounter++
					log.Infof("Deleting torrent file: %s", fc.GetLastDownloadedFileName())
					err = os.Remove(fc.GetLastDownloadedFileName())
					if err != nil {
						log.Error(err)
					}
					err = qc.DeleteTorrents([]string{t.Hash})
					if err != nil {
						log.Error(err)
					}
				}
			}

		}
		torrents, err = qc.GetAllTorrents(nil)
		if err != nil {
			log.Error(err)
		}
		log.Infof("Torrents count: %d [Updated %d torrents]", len(torrents), updateCounter)
	}

}
