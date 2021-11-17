// Butterfly - The web crawler base on Apache Solr for StarStart!
// Copyright(c) 2020 Star Inc. All Rights Reserved.
// The software licensed under Mozilla Public License Version 2.0

package http

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
	"strings"
)

type StatusCode int

const (
	DefaultUserAgent = "Mozilla/5.0 (compatible; Star Butterfly/1.0; +https://github.com/star-inc/butterfly)"
)

type Client struct {
	baseURL   *url.URL
	userAgent string
}

func NewHttpClient(baseURL string) *Client {
	httpClient := new(Client)
	httpClient.SetBaseURL(baseURL)
	return httpClient
}

func (c *Client) SetBaseURL(baseURL string) *Client {
	var err error
	if c.baseURL, err = url.Parse(baseURL); err != nil {
		log.Panicln(err)
	}
	return c
}

func (c *Client) SetUserAgent(userAgent string) *Client {
	c.userAgent = userAgent
	return c
}

func (c *Client) baseURLGlue(uri string) string {
	if strings.HasPrefix(uri, "http://") ||
		strings.HasPrefix(uri, "https://") {
		return uri
	} else {
		newURL := *c.baseURL
		newURL.Path = path.Join(newURL.Path, uri)
		return newURL.String()
	}
}

func (c *Client) initRequest(method, uri string, body io.Reader) *http.Request {
	request, err := http.NewRequest(method, uri, body)
	if err != nil {
		log.Panicln(err)
	}
	if c.userAgent != "" {
		request.Header.Add("User-Agent", c.userAgent)
	} else {
		request.Header.Add("User-Agent", DefaultUserAgent)
	}
	return request
}

func (c *Client) Do(method, uri string, data interface{}) (StatusCode, []byte) {
	client := &http.Client{}
	fullURI := c.baseURLGlue(uri)
	var request *http.Request
	if reader, ok := data.(io.Reader); ok {
		request = c.initRequest(method, fullURI, reader)
	} else {
		request = c.initRequest(method, fullURI, nil)
	}
	if _, ok := data.(bytes.Buffer); ok {
		request.Header.Add("Content-Type", "application/json; charset=utf-8")
	} else if _, ok := data.(strings.Reader); ok {
		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}
	response, err := client.Do(request)
	if err != nil {
		log.Panicln(err)
	}
	defer func() {
		if err := response.Body.Close(); err != nil {
			log.Panicln(err)
		}
	}()
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Panicln(err)
	}
	return StatusCode(response.StatusCode), result
}

func (c *Client) GET(uri string) (StatusCode, []byte) {
	return c.Do("GET", uri, nil)
}

func (c *Client) POST(uri string, data interface{}) (StatusCode, []byte) {
	return c.Do("POST", uri, data)
}

func (c *Client) PUT(uri string, data interface{}) (StatusCode, []byte) {
	return c.Do("PUT", uri, data)
}

func (c *Client) DELETE(uri string, data interface{}) (StatusCode, []byte) {
	return c.Do("DELETE", uri, data)
}

func (c *Client) PATCH(uri string, data interface{}) (StatusCode, []byte) {
	return c.Do("PATCH", uri, data)
}
