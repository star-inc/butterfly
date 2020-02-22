/*
Package butterfly : The library for butterfly

Copyright(c) 2020 Star Inc. All Rights Reserved.
This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/
package butterfly

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/gocolly/colly"
	strip "github.com/grokify/html-strip-tags-go"
)

// CollyHandle :
type CollyHandle struct {
	Client    *colly.Collector
	UserAgent string
	Content   string
}

func removeSyntaxs(rawString string) string {
	rawString = strings.ReplaceAll(rawString, " ", "")
	rawString = strings.ReplaceAll(rawString, "\t", "")
	rawString = strings.ReplaceAll(rawString, "\n", "")
	return rawString
}

// NewCollyClient :
func NewCollyClient(userAgent string) *CollyHandle {
	handle := new(CollyHandle)
	handle.setUserAgent(userAgent)
	client := colly.NewCollector(
		colly.UserAgent(handle.UserAgent),
	)
	handle.Client = client
	return handle
}

func (handle *CollyHandle) setUserAgent(userAgent string) {
	if userAgent == "" {
		handle.UserAgent = "Mozilla/5.0 (compatible; Star Butterfly/1.0; +https://github.com/star-inc/butterfly)"
	} else {
		handle.UserAgent = userAgent
	}
}

// Fetch :
func (handle *CollyHandle) Fetch(uri string, solrHandle *SolrHandle) {
	data := new(VioletDataStruct)
	url, _ := url.Parse(uri)
	handle.Client.Async = true
	handle.Client.AllowedDomains = []string{url.Host}
	handle.Client.OnHTML("title", func(e *colly.HTMLElement) {
		data.Title = e.Text
	})
	handle.Client.OnHTML("div", func(e *colly.HTMLElement) {
		data.Description = removeSyntaxs(strip.StripTags(e.Text))
	})
	handle.Client.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})
	handle.Client.OnRequest(func(r *colly.Request) {
		data.URI = r.URL.String()
		fmt.Println("Visiting", r.URL)
		solrHandle.Update(data)
	})
	handle.Client.Visit(uri)
}
