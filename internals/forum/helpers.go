package forum

import (
	"bytes"
	"context"
	"fmt"
	"golang.org/x/net/proxy"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
)

func (client *ForumClient) post(uri string, opts map[string]string, body io.Reader, contentType string) (*http.Response, error) {
	client.log.Debugf("POST URL: %s", client.url+uri)
	var b strings.Builder
	if body == nil {
		for k, v := range opts {
			s := fmt.Sprintf("%s=%s&", k, v)
			b.WriteString(s)
		}
		body = bytes.NewBufferString(b.String())
	}
	req, err := http.NewRequest(http.MethodPost, client.url+uri, body)
	if err != nil {
		return nil, err
	}
	for k, v := range opts {
		client.log.Debugf("opts: %s: %s", k, v)
	}
	req.Proto = "HTTP/2.0"
	req.ProtoMajor = 2
	req.ProtoMinor = 0
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/109.0")
	if contentType == "" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=windows-1251")
	} else {
		req.Header.Set("Content-Type", contentType)
	}

	for k, v := range req.Header {
		client.log.Debugf("request: %s: %s", k, v)
	}
	resp, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	for k, v := range resp.Header {
		client.log.Debugf("response: %s: %s", k, v)
	}
	return resp, nil
}

func (client *ForumClient) get(uri string, opts map[string]string) (*http.Response, error) {
	client.log.Debugf("GET URL: %s", client.url+uri)
	req, err := http.NewRequest(http.MethodGet, client.url+uri, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/109.0")
	if opts != nil {
		query := req.URL.Query()
		for k, v := range opts {
			query.Add(k, v)
		}
		req.URL.RawQuery = query.Encode()
	}
	client.log.Debug(req.URL)
	client.log.Debug(req.Header)
	resp, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (client *ForumClient) saveTorrentToFile(fName string, body []byte) error {
	const downloadDir = "downloads"
	path := fmt.Sprintf("%s/%s.torrent", downloadDir, fName)
	if _, err := os.Stat(downloadDir); os.IsNotExist(err) {
		if err := os.Mkdir(downloadDir, os.ModePerm); err != nil {
			return err
		}
		return err
	}
	client.log.Infof("Saving torrent file to: %s", path)
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(body)
	if err != nil {
		return err
	}
	client.lastFile = path
	return nil
}

func (client *ForumClient) configureProxy() error {
	dialer, err := proxy.SOCKS5("tcp", client.socks5, nil, proxy.Direct)
	if err != nil {
		return err
	}
	//ctx := context.TODO()
	dialContext := func(ctx context.Context, network, address string) (net.Conn, error) {
		return dialer.Dial(network, address)
	}
	transport := &http.Transport{DialContext: dialContext,
		DisableKeepAlives: true}
	client.httpClient = &http.Client{Transport: transport}
	return nil
}
