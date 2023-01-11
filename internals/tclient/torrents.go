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
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func (c *qClient) GetAllTorrents(filters map[string]string) ([]BasicTorrent, error) {
	resp, err := c.get("/api/v2/torrents/info", filters)
	if err != nil {
		return nil, err
	}
	var torrentList []BasicTorrent
	err = json.NewDecoder(resp.Body).Decode(&torrentList)
	if err != nil {
		return nil, err
	}
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
	err = json.NewDecoder(resp.Body).Decode(&torrent)
	if err != nil {
		return nil, err
	}
	return &torrent, nil
}
