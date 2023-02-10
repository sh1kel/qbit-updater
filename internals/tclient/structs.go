package tclient

type BasicTorrent struct {
	AddedOn       int     `json:"added_on"`
	Category      string  `json:"category"`
	CompletionOn  int64   `json:"completion_on"`
	Dlspeed       int     `json:"dlspeed"`
	Eta           int     `json:"eta"`
	ForceStart    bool    `json:"force_start"`
	Hash          string  `json:"hash"`
	Name          string  `json:"name"`
	NumComplete   int     `json:"num_complete"`
	NumIncomplete int     `json:"num_incomplete"`
	NumLeechs     int     `json:"num_leechs"`
	NumSeeds      int     `json:"num_seeds"`
	Priority      int     `json:"priority"`
	Progress      float64 `json:"progress"`
	Ratio         float64 `json:"ratio"`
	SavePath      string  `json:"save_path"`
	SeqDl         bool    `json:"seq_dl"`
	Size          int     `json:"size"`
	State         string  `json:"state"`
	SuperSeeding  bool    `json:"super_seeding"`
	Upspeed       int     `json:"upspeed"`
}

type Torrent struct {
	AdditionDate           int     `json:"addition_date"`
	Comment                string  `json:"comment"`
	CompletionDate         int     `json:"completion_date"`
	CreatedBy              string  `json:"created_by"`
	CreationDate           int     `json:"creation_date"`
	DlLimit                int     `json:"dl_limit"`
	DlSpeed                int     `json:"dl_speed"`
	DlSpeedAvg             int     `json:"dl_speed_avg"`
	Eta                    int     `json:"eta"`
	LastSeen               int     `json:"last_seen"`
	NbConnections          int     `json:"nb_connections"`
	NbConnectionsLimit     int     `json:"nb_connections_limit"`
	Peers                  int     `json:"peers"`
	PeersTotal             int     `json:"peers_total"`
	PieceSize              int     `json:"piece_size"`
	PiecesHave             int     `json:"pieces_have"`
	PiecesNum              int     `json:"pieces_num"`
	Reannounce             int     `json:"reannounce"`
	SavePath               string  `json:"save_path"`
	SeedingTime            int     `json:"seeding_time"`
	Seeds                  int     `json:"seeds"`
	SeedsTotal             int     `json:"seeds_total"`
	ShareRatio             float64 `json:"share_ratio"`
	TimeElapsed            int     `json:"time_elapsed"`
	TotalDownloaded        int     `json:"total_downloaded"`
	TotalDownloadedSession int     `json:"total_downloaded_session"`
	TotalSize              int     `json:"total_size"`
	TotalUploaded          int     `json:"total_uploaded"`
	TotalUploadedSession   int     `json:"total_uploaded_session"`
	TotalWasted            int     `json:"total_wasted"`
	UpLimit                int     `json:"up_limit"`
	UpSpeed                int     `json:"up_speed"`
	UpSpeedAvg             int     `json:"up_speed_avg"`
}

type Tracker struct {
	Msg      string `json:"msg"`
	NumPeers int    `json:"num_peers"`
	Status   int    `json:"status"`
	URL      string `json:"url"`
}

type TorrentFile struct {
	IsSeed   bool   `json:"is_seed"`
	Name     string `json:"name"`
	Priority int    `json:"priority"`
	Progress int    `json:"progress"`
	Size     int    `json:"size"`
}

type QbConfig struct {
	AddTrackers                        string `json:"add_trackers"`
	AddTrackersEnabled                 bool   `json:"add_trackers_enabled"`
	AltDlLimit                         int64  `json:"alt_dl_limit"`
	AltUpLimit                         int64  `json:"alt_up_limit"`
	AlternativeWebuiEnabled            bool   `json:"alternative_webui_enabled"`
	AlternativeWebuiPath               string `json:"alternative_webui_path"`
	AnnounceIP                         string `json:"announce_ip"`
	AnnounceToAllTiers                 bool   `json:"announce_to_all_tiers"`
	AnnounceToAllTrackers              bool   `json:"announce_to_all_trackers"`
	AnonymousMode                      bool   `json:"anonymous_mode"`
	AsyncIoThreads                     int64  `json:"async_io_threads"`
	AutoDeleteMode                     int64  `json:"auto_delete_mode"`
	AutoTmmEnabled                     bool   `json:"auto_tmm_enabled"`
	AutorunEnabled                     bool   `json:"autorun_enabled"`
	AutorunProgram                     string `json:"autorun_program"`
	BannedIPs                          string `json:"banned_IPs"`
	BittorrentProtocol                 int64  `json:"bittorrent_protocol"`
	BypassAuthSubnetWhitelist          string `json:"bypass_auth_subnet_whitelist"`
	BypassAuthSubnetWhitelistEnabled   bool   `json:"bypass_auth_subnet_whitelist_enabled"`
	BypassLocalAuth                    bool   `json:"bypass_local_auth"`
	CategoryChangedTmmEnabled          bool   `json:"category_changed_tmm_enabled"`
	CheckingMemoryUse                  int64  `json:"checking_memory_use"`
	CreateSubfolderEnabled             bool   `json:"create_subfolder_enabled"`
	CurrentInterfaceAddress            string `json:"current_interface_address"`
	CurrentNetworkInterface            string `json:"current_network_interface"`
	Dht                                bool   `json:"dht"`
	DiskCache                          int64  `json:"disk_cache"`
	DiskCacheTTL                       int64  `json:"disk_cache_ttl"`
	DlLimit                            int64  `json:"dl_limit"`
	DontCountSlowTorrents              bool   `json:"dont_count_slow_torrents"`
	DyndnsDomain                       string `json:"dyndns_domain"`
	DyndnsEnabled                      bool   `json:"dyndns_enabled"`
	DyndnsPassword                     string `json:"dyndns_password"`
	DyndnsService                      int64  `json:"dyndns_service"`
	DyndnsUsername                     string `json:"dyndns_username"`
	EmbeddedTrackerPort                int64  `json:"embedded_tracker_port"`
	EnableCoalesceReadWrite            bool   `json:"enable_coalesce_read_write"`
	EnableEmbeddedTracker              bool   `json:"enable_embedded_tracker"`
	EnableMultiConnectionsFromSameIP   bool   `json:"enable_multi_connections_from_same_ip"`
	EnableOsCache                      bool   `json:"enable_os_cache"`
	EnablePieceExtentAffinity          bool   `json:"enable_piece_extent_affinity"`
	EnableUploadSuggestions            bool   `json:"enable_upload_suggestions"`
	Encryption                         int64  `json:"encryption"`
	ExportDir                          string `json:"export_dir"`
	ExportDirFin                       string `json:"export_dir_fin"`
	FilePoolSize                       int64  `json:"file_pool_size"`
	IncompleteFilesExt                 bool   `json:"incomplete_files_ext"`
	IPFilterEnabled                    bool   `json:"ip_filter_enabled"`
	IPFilterPath                       string `json:"ip_filter_path"`
	IPFilterTrackers                   bool   `json:"ip_filter_trackers"`
	LimitLanPeers                      bool   `json:"limit_lan_peers"`
	LimitTcpOverhead                   bool   `json:"limit_tcp_overhead"`
	LimitUtpRate                       bool   `json:"limit_utp_rate"`
	ListenPort                         int64  `json:"listen_port"`
	Locale                             string `json:"locale"`
	Lsd                                bool   `json:"lsd"`
	MailNotificationAuthEnabled        bool   `json:"mail_notification_auth_enabled"`
	MailNotificationEmail              string `json:"mail_notification_email"`
	MailNotificationEnabled            bool   `json:"mail_notification_enabled"`
	MailNotificationPassword           string `json:"mail_notification_password"`
	MailNotificationSender             string `json:"mail_notification_sender"`
	MailNotificationSMTP               string `json:"mail_notification_smtp"`
	MailNotificationSslEnabled         bool   `json:"mail_notification_ssl_enabled"`
	MailNotificationUsername           string `json:"mail_notification_username"`
	MaxActiveDownloads                 int64  `json:"max_active_downloads"`
	MaxActiveTorrents                  int64  `json:"max_active_torrents"`
	MaxActiveUploads                   int64  `json:"max_active_uploads"`
	MaxConnec                          int64  `json:"max_connec"`
	MaxConnecPerTorrent                int64  `json:"max_connec_per_torrent"`
	MaxRatio                           int64  `json:"max_ratio"`
	MaxRatioAct                        int64  `json:"max_ratio_act"`
	MaxRatioEnabled                    bool   `json:"max_ratio_enabled"`
	MaxSeedingTime                     int64  `json:"max_seeding_time"`
	MaxSeedingTimeEnabled              bool   `json:"max_seeding_time_enabled"`
	MaxUploads                         int64  `json:"max_uploads"`
	MaxUploadsPerTorrent               int64  `json:"max_uploads_per_torrent"`
	OutgoingPortsMax                   int64  `json:"outgoing_ports_max"`
	OutgoingPortsMin                   int64  `json:"outgoing_ports_min"`
	Pex                                bool   `json:"pex"`
	PreallocateAll                     bool   `json:"preallocate_all"`
	ProxyAuthEnabled                   bool   `json:"proxy_auth_enabled"`
	ProxyIP                            string `json:"proxy_ip"`
	ProxyPassword                      string `json:"proxy_password"`
	ProxyPeerConnections               bool   `json:"proxy_peer_connections"`
	ProxyPort                          int64  `json:"proxy_port"`
	ProxyTorrentsOnly                  bool   `json:"proxy_torrents_only"`
	ProxyType                          int64  `json:"proxy_type"`
	ProxyUsername                      string `json:"proxy_username"`
	QueueingEnabled                    bool   `json:"queueing_enabled"`
	RandomPort                         bool   `json:"random_port"`
	RecheckCompletedTorrents           bool   `json:"recheck_completed_torrents"`
	ResolvePeerCountries               bool   `json:"resolve_peer_countries"`
	RssAutoDownloadingEnabled          bool   `json:"rss_auto_downloading_enabled"`
	RssDownloadRepackProperEpisodes    bool   `json:"rss_download_repack_proper_episodes"`
	RssMaxArticlesPerFeed              int64  `json:"rss_max_articles_per_feed"`
	RssProcessingEnabled               bool   `json:"rss_processing_enabled"`
	RssRefreshInterval                 int64  `json:"rss_refresh_interval"`
	RssSmartEpisodeFilters             string `json:"rss_smart_episode_filters"`
	SavePath                           string `json:"save_path"`
	SavePathChangedTmmEnabled          bool   `json:"save_path_changed_tmm_enabled"`
	SaveResumeDataInterval             int64  `json:"save_resume_data_interval"`
	ScheduleFromHour                   int64  `json:"schedule_from_hour"`
	ScheduleFromMin                    int64  `json:"schedule_from_min"`
	ScheduleToHour                     int64  `json:"schedule_to_hour"`
	ScheduleToMin                      int64  `json:"schedule_to_min"`
	SchedulerDays                      int64  `json:"scheduler_days"`
	SchedulerEnabled                   bool   `json:"scheduler_enabled"`
	SendBufferLowWatermark             int64  `json:"send_buffer_low_watermark"`
	SendBufferWatermark                int64  `json:"send_buffer_watermark"`
	SendBufferWatermarkFactor          int64  `json:"send_buffer_watermark_factor"`
	SlowTorrentDlRateThreshold         int64  `json:"slow_torrent_dl_rate_threshold"`
	SlowTorrentInactiveTimer           int64  `json:"slow_torrent_inactive_timer"`
	SlowTorrentUlRateThreshold         int64  `json:"slow_torrent_ul_rate_threshold"`
	SocketBacklogSize                  int64  `json:"socket_backlog_size"`
	StartPausedEnabled                 bool   `json:"start_paused_enabled"`
	StopTrackerTimeout                 int64  `json:"stop_tracker_timeout"`
	TempPath                           string `json:"temp_path"`
	TempPathEnabled                    bool   `json:"temp_path_enabled"`
	TorrentChangedTmmEnabled           bool   `json:"torrent_changed_tmm_enabled"`
	UpLimit                            int64  `json:"up_limit"`
	UploadChokingAlgorithm             int64  `json:"upload_choking_algorithm"`
	UploadSlotsBehavior                int64  `json:"upload_slots_behavior"`
	Upnp                               bool   `json:"upnp"`
	UseHTTPS                           bool   `json:"use_https"`
	UtpTcpMixedMode                    int64  `json:"utp_tcp_mixed_mode"`
	WebUIAddress                       string `json:"web_ui_address"`
	WebUIBanDuration                   int64  `json:"web_ui_ban_duration"`
	WebUIClickjackingProtectionEnabled bool   `json:"web_ui_clickjacking_protection_enabled"`
	WebUICsrfProtectionEnabled         bool   `json:"web_ui_csrf_protection_enabled"`
	WebUICustomHTTPHeaders             string `json:"web_ui_custom_http_headers"`
	WebUIDomainList                    string `json:"web_ui_domain_list"`
	WebUIHostHeaderValidationEnabled   bool   `json:"web_ui_host_header_validation_enabled"`
	WebUIHTTPSCertPath                 string `json:"web_ui_https_cert_path"`
	WebUIHTTPSKeyPath                  string `json:"web_ui_https_key_path"`
	WebUIMaxAuthFailCount              int64  `json:"web_ui_max_auth_fail_count"`
	WebUIPort                          int64  `json:"web_ui_port"`
	WebUISecureCookieEnabled           bool   `json:"web_ui_secure_cookie_enabled"`
	WebUISessionTimeout                int64  `json:"web_ui_session_timeout"`
	WebUIUpnp                          bool   `json:"web_ui_upnp"`
	WebUIUseCustomHTTPHeadersEnabled   bool   `json:"web_ui_use_custom_http_headers_enabled"`
	WebUIUsername                      string `json:"web_ui_username"`
}

const (
	TrackerIsDisabled                        int    = 0
	TrackerHasNotBeenContactedYet            int    = 1
	TrackerHasBeenContactedAndIsWorking      int    = 2
	TrackerIsUpdating                        int    = 3
	TrackerHasBeenContactedButItIsNotWorking int    = 4
	TorrentNotRegistered                     string = "Torrent not registered"
)
