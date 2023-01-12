package tclient

import (
	"bytes"
	"fmt"
	"html"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"net/url"
	"os"
)

func (c *qClient) post(uri string, opts map[string]string, body io.Reader, contentType string) (*http.Response, error) {
	c.log.Debugf("POST URL: %s", c.url+uri)
	req, err := http.NewRequest(http.MethodPost, c.url+uri, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "go-qbittorrent v0.1")
	if contentType == "" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	} else {
		req.Header.Set("Content-Type", contentType)
	}
	if opts != nil {
		form := url.Values{}
		for k, v := range opts {
			form.Add(k, v)
		}
		req.PostForm = form
	}
	c.log.Debugf("%v", req)
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
	c.log.Debug(req.URL)
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *qClient) createReqBody(file string, options map[string]string) (io.Reader, string, error) {
	fileBuffer := new(bytes.Buffer)
	writer := multipart.NewWriter(fileBuffer)
	tFile, err := os.Open(file)
	if err != nil {
		return nil, "", err
	}

	defer tFile.Close()

	for key, val := range options {
		err = writer.WriteField(key, val)
		if err != nil {
			return nil, "", err
		}
	}

	header := make(textproto.MIMEHeader)
	header.Set("Content-Disposition", fmt.Sprintf(`form-data; name="torrents"; filename="%s"`, html.EscapeString(file)))
	header.Set("Content-Type", "application/x-bittorrent")

	formWriter, err := writer.CreatePart(header)
	if err != nil {
		return nil, "", err
	}
	_, err = io.Copy(formWriter, tFile)
	if err != nil {
		return nil, "", err
	}

	err = writer.Close()
	if err != nil {
		return nil, "", err
	}

	return fileBuffer, writer.FormDataContentType(), nil
}
