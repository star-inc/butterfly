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

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

// CollyHandle : Define CollyHandle Class
type CollyHandle struct {
	Client    *colly.Collector
	UserAgent string
}

// NewCollyClient : Create new colly collection
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
		handle.UserAgent = ""
	} else {
		handle.UserAgent = userAgent
	}
}

// Fetch : Capture web pages on Internet and submit to Solr
func (handle *CollyHandle) Fetch(uri string, solrHandle *SolrHandle) {
	data := new(VioletDataStruct)
	url, _ := url.Parse(uri)
	colly.Async(true)
	handle.Client.AllowedDomains = []string{url.Host}

	handle.Client.OnHTML("meta[name=description]", func(e *colly.HTMLElement) {
		data.Description = e.Attr("content")
	})

	handle.Client.OnHTML("title", func(e *colly.HTMLElement) {
		data.Title = e.Text
	})

	handle.Client.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	handle.Client.OnRequest(func(r *colly.Request) {
		data.URI = r.URL.String()
		fmt.Println("Visiting", r.URL)

		capturedHTML := HTTPGet(data.URI)
		reader := strings.NewReader(capturedHTML)
		doc, err := goquery.NewDocumentFromReader(reader)
		DeBug("Load HTML", err)
		doc.Find("script").Remove() // Remove Javascript codes
		doc.Find("style").Remove()  // Remove CSS codes
		data.Content = ReplaceSyntaxs(doc.Text(), " ")

		solrHandle.Update(data)
	})

	handle.Client.Visit(uri)
	handle.Client.Wait()
}
