package tclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"
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

func (c *qClient) DownloadFromFile(file string, options map[string]string) error {
	c.log.Debugf("Add torrent from file: %s", file)
	reader, ct, err := c.createReqBody(file, options)
	if err != nil {
		return err
	}

	resp, err := c.post("/api/v2/torrents/add", options, reader, ct)
	if err != nil {
		return err
	}
	resp.Body.Close()
	return nil
}

func (c *qClient) DeleteTorrents(hashes []string) error {
	c.log.Infof("Deleting torrent: %v", hashes)

	hashList := strings.Builder{}
	for i, hash := range hashes {
		if i < len(hashes)-1 {
			hashList.WriteString(fmt.Sprintf("%s|", hash))
		} else {
			hashList.WriteString(hash)
		}
	}
	options := map[string]string{
		"hashes":      hashList.String(),
		"deleteFiles": "false",
	}
	var b strings.Builder
	for k, v := range options {
		s := fmt.Sprintf("%s=%s&", k, v)
		b.WriteString(s)
	}
	//body := bytes.NewBufferString(url.QueryEscape(b.String()))
	body := bytes.NewBufferString(b.String())

	resp, err := c.post("/api/v2/torrents/delete", options, body, "")
	//resp, err := c.get("/api/v2/torrents/delete", options)
	if err != nil {
		return err
	}
	c.log.Infof("Status: %s", resp.Status)
	return nil
}

func (c *qClient) GetTrackers(hash string) ([]Tracker, error) {
	var t []Tracker
	opts := map[string]string{"hash": hash}

	resp, err := c.get("/api/v2/torrents/trackers", opts)
	if err != nil {
		return t, err
	}
	json.NewDecoder(resp.Body).Decode(&t)
	return t, nil
}

func (c *qClient) GetShortHashFromComment(hash string) (string, error) {
	t, err := c.GetTorrentInfo(hash)
	if err != nil {
		return "", err
	}
	s := strings.Split(t.Comment, "=")
	if len(s) < 2 {
		return "", fmt.Errorf("bad comment: %s", t.Comment)
	}
	return s[len(s)-1], nil
}
