package tclient

import (
	"encoding/json"
	"io"
)

func (c *qClient) GetVersion() (string, error) {
	resp, err := c.get("/api/v2/app/version", nil)
	if err != nil {
		return "", err
	}
	body, err := io.ReadAll(resp.Body)
	return string(body), nil
}

func (c *qClient) GetAllTorrents(filters map[string]string) ([]BasicTorrent, error) {
	resp, err := c.get("/api/v2/torrents/info", filters)
	if err != nil {
		return nil, err
	}
	var torrentList []BasicTorrent
	//body, err := io.ReadAll(resp.Body)
	json.NewDecoder(resp.Body).Decode(&torrentList)
	return torrentList, nil
}

func (c *qClient) GetTorrentInfo(hash string) (*Torrent, error) {
	c.log.Debug(hash)
	opts := map[string]string{"hash": hash}
	resp, err := c.get("/api/v2/torrents/properties", opts)
	if err != nil {
		return nil, err
	}
	var torrent Torrent
	json.NewDecoder(resp.Body).Decode(&torrent)
	return &torrent, nil
}
