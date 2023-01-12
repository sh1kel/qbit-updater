package tclient

import (
	"errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/publicsuffix"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

type qClient struct {
	url           string
	user          string
	pass          string
	client        *http.Client
	jar           http.CookieJar
	log           *logrus.Logger
	authenticated bool
}

func New(url string, user string, pass string, log *logrus.Logger) *qClient {
	jar, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	return &qClient{
		url:           url,
		user:          user,
		pass:          pass,
		client:        new(http.Client),
		jar:           jar,
		log:           log,
		authenticated: false,
	}

}

func (c *qClient) Connect() error {
	credentials := make(map[string]string)
	credentials["username"] = c.user
	credentials["password"] = c.pass
	resp, err := c.post("/api/v2/auth/login", credentials, nil, "")
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New("got non 200 response code")
	}
	c.log.Debugf("Response code: %d", resp.StatusCode)
	if cookies := resp.Cookies(); len(cookies) > 0 {
		cookieURL, _ := url.Parse(c.url)
		c.jar.SetCookies(cookieURL, cookies)
	}
	c.client = &http.Client{
		Jar: c.jar,
	}
	return nil
}
