package forum

import (
	"fmt"
	"github.com/sh1kel/qbit-updater/internals/configuration"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/publicsuffix"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
)

type ForumClient struct {
	userName   string
	pass       string
	login      string
	url        string
	httpClient *http.Client
	log        *logrus.Logger
	socks5     string
	useProxy   bool
	lastFile   string
}

func New(config *configuration.Config) *ForumClient {
	return &ForumClient{
		userName: config.Forum.UserName,
		pass:     config.Forum.UserPass,
		login:    config.Forum.Login,
		url:      config.Forum.Url,
		log:      config.Logger,
		socks5:   config.Forum.Socks5,
		useProxy: config.Forum.UseProxy,
	}
}

func (client *ForumClient) Init() error {
	if client.useProxy {
		client.log.Infof("Use socks proxy: %s", client.socks5)
		err := client.configureProxy()
		if err != nil {
			return err
		}
		client.httpClient.Jar, _ = cookiejar.New(&cookiejar.Options{
			PublicSuffixList: publicsuffix.List,
		})
		return nil
	}
	client.httpClient = &http.Client{}
	client.httpClient.Jar, _ = cookiejar.New(&cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	})
	return nil
}

func (client *ForumClient) Auth() error {
	opts := map[string]string{
		"login_username": url.QueryEscape(client.userName),
		"login_password": url.QueryEscape(client.pass),
		"login":          client.login,
	}
	resp, err := client.post("/forum/login.php", opts, nil, "")
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("got error response code: %d", resp.StatusCode)
	}
	client.log.Debugf("Response code: %d", resp.StatusCode)
	client.log.Debugf("Cookies: %v", resp.Cookies())

	if cookies := resp.Cookies(); len(cookies) > 0 {
		cookieURL, _ := url.Parse(client.url)
		client.httpClient.Jar.SetCookies(cookieURL, cookies)
	}
	return nil
}

func (client *ForumClient) GetTorrentFile(torrentId string) error {
	opts := map[string]string{
		"t": torrentId,
	}
	resp, err := client.get("/forum/dl.php", opts)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("got response code: %d", resp.StatusCode)
	}
	client.log.Debugf("Response code: %d", resp.StatusCode)
	client.log.Debugf("Content type: %s", resp.Cookies())
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = client.saveTorrentToFile(torrentId, body)
	if err != nil {
		return err
	}
	return nil
}

func (client *ForumClient) CleanFile() error {
	err := os.Remove(client.lastFile)
	if err != nil {
		return err
	}
	return nil
}
