package tclient

import (
	"net/http"
	"net/url"
)

func (c *qClient) post(uri string, opts map[string]string) (*http.Response, error) {
	c.log.Debugf("POST URL: %s", c.url+uri)
	req, err := http.NewRequest(http.MethodPost, c.url+uri, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "go-qbittorrent v0.1")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if opts != nil {
		form := url.Values{}
		for k, v := range opts {
			form.Add(k, v)
		}
		req.PostForm = form
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *qClient) get(uri string, opts map[string]string) (*http.Response, error) {
	c.log.Debugf("GET URL: %s", c.url+uri)
	req, err := http.NewRequest(http.MethodGet, c.url+uri, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "go-qbittorrent v0.1")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if opts != nil {
		query := req.URL.Query()
		for k, v := range opts {
			query.Add(k, v)
		}
		req.URL.RawQuery = query.Encode()
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
