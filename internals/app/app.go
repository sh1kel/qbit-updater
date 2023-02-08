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
		qc := tclient.New(url, "1", "2", log)
		err = qc.Connect()
		if err != nil {
			log.Error(err)
		}
		version, _ := qc.GetVersion()
		log.Infof("qB version: %s", version)
		if err != nil {
			log.Error(err)
		}
		torrents, err := qc.GetAllTorrents(nil)
		if err != nil {
			log.Error(err)
		}
		for _, t := range torrents {
			tt, _ := qc.GetTrackers(t.Hash)
			for _, tracker := range tt {
				if tracker.Status == 4 {
					log.Infof("Status: %d", tracker.Status)
					ti, _ := qc.GetTorrentInfo(t.Hash)
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
						err = os.Remove(fc.GetLastDownloadedFileName())
						if err != nil {
							log.Error(err)
						}
					}
					qc.DeleteTorrents([]string{t.Hash})

				}
			}

		}
	}

}
