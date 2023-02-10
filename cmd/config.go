package cmd

import (
	"github.com/sh1kel/qbit-updater/internals/app"
	"github.com/sh1kel/qbit-updater/internals/configuration"
	"github.com/spf13/cobra"
)

var (
	configCmd = &cobra.Command{
		Use:   "config",
		Short: "set qbittorrent config options",
		Long:  ``,
		Run:   start,
	}
	// Speed tab: speed limits
	dlLimit string // Upload
	upLimit string // Download
	// Connection tab: connection limits
	maxConnections           string // Global maximum number of connections
	maxConnectionsPerTorrent string // Maximum number of connections per torrent
	maxUploads               string // Global maximum number of upload slots
	maxUploadsPerTorrent     string // Maximum number of upload slot per torrent
	// BitTorrent tab: torrent queueing
	maxActiveDownloads string // Maximum active downloads
	maxActiveUploads   string // Maximum active uploads
	maxActiveTorrents  string // Maximum active torrents

	params = make(map[string]string)
)

func init() {
	rootCmd.AddCommand(configCmd)
	rootCmd.Flags().StringVar(&dlLimit, "dl-limit", "0", "")
	rootCmd.Flags().StringVar(&upLimit, "up-limit", "6291456", "")
	rootCmd.Flags().StringVar(&maxConnections, "max-connec", "5000", "")
	rootCmd.Flags().StringVar(&maxConnectionsPerTorrent, "max-connec-per-torrent", "2", "")
	rootCmd.Flags().StringVar(&maxUploads, "max-uploads", "5", "")
	rootCmd.Flags().StringVar(&maxUploadsPerTorrent, "max-uploads-per-torrent", "10", "")
	rootCmd.Flags().StringVar(&maxActiveDownloads, "max-active-downloads", "20", "")
	rootCmd.Flags().StringVar(&maxActiveUploads, "max-active-uploads", "5", "")
	rootCmd.Flags().StringVar(&maxActiveTorrents, "max-active-torrents", "10000", "")

	initConfigMap()
}

func start(cmd *cobra.Command, args []string) {
	config := configuration.InitConfig(configFile)
	err := configureLogger(config)
	if err != nil {
		config.Logger.Fatal(err)
	}
	app.SetConfigOptions(params, config)
}

func initConfigMap() {
	params["dl_limit"] = dlLimit
	params["up_limit"] = upLimit
	params["max_connec"] = maxConnections
	params["max_connec_per_torrent"] = maxConnectionsPerTorrent
	params["max_uploads"] = maxUploads
	params["max_uploads_per_torrent"] = maxUploadsPerTorrent
	params["max_active_downloads"] = maxActiveDownloads
	params["max_active_uploads"] = maxActiveUploads
	params["max_active_torrents"] = maxActiveTorrents
}
