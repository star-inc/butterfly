/*
Package butterfly : The library for butterfly

Copyright(c) 2020 Star Inc. All Rights Reserved.
This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/
package butterfly

import (
	"crypto/md5"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/queue"
	"github.com/temoto/robotstxt"
	"github.com/velebak/colly-sqlite3-storage/colly/sqlite3"
)

// Handles : Define Handles Class
type Handles struct {
	Colly        *colly.Collector
	CollyStorage *sqlite3.Storage
	Solr         *SolrHandle
	RobotsTXT    map[string]string
}

// NewCollyClient : Create new colly collection
func NewCollyClient(solr *SolrHandle) *Handles {
	handle := new(Handles)
	client := colly.NewCollector(
		colly.UserAgent(Config.UserAgent),
	)
	handle.Colly = client
	handle.Solr = solr
	handle.RobotsTXT = make(map[string]string)
	return handle
}

func (handle *Handles) setStorage(domain string) {
	if _, err := os.Stat(Config.Colly.SqlitePath); os.IsNotExist(err) {
		err = os.MkdirAll(Config.Colly.SqlitePath, 0755)
		DeBug("Colly setStorage create directory", err)
	}
	path := fmt.Sprintf("%s%s.sqlite3", Config.Colly.SqlitePath, domain)
	storage := &sqlite3.Storage{
		Filename: path,
	}
	err := handle.Colly.SetStorage(storage)
	DeBug("Colly setStorage", err)
	handle.CollyStorage = storage
}

// Fetch : Capture web pages on Internet and submit to Solr
func (handle *Handles) Fetch(uri string) {
	data := new(VioletDataStruct)

	queryURL, _ := url.Parse(uri)
	handle.setStorage(queryURL.Host)
	defer handle.CollyStorage.Close()

	if Config.ForcusOnURI {
		handle.Colly.AllowedDomains = []string{queryURL.Host}
	}

	var collyQueue *queue.Queue
	collyQueue, _ = queue.New(Config.Colly.Threads, handle.CollyStorage)

	handle.Colly.OnHTML("meta[name=description]", func(e *colly.HTMLElement) {
		data.Description = e.Attr("content")
	})

	handle.Colly.OnHTML("title", func(e *colly.HTMLElement) {
		data.Title = e.Text
	})

	handle.Colly.OnHTML("a[href]", func(e *colly.HTMLElement) {
		collyQueue.AddURL(e.Request.AbsoluteURL(e.Attr("href")))
	})

	handle.Colly.OnRequest(func(r *colly.Request) {
		data.URI = r.URL.String()
		fmt.Println("Visiting", r.URL)

		signature := md5.Sum([]byte(data.URI))
		data.ID = fmt.Sprintf("%x", signature)

		if _, exists := handle.RobotsTXT[r.URL.Host]; !exists {
			handle.RobotsTXT[r.URL.Host] = HTTPGet(fmt.Sprintf("%s://%s/robots.txt", r.URL.Scheme, r.URL.Host), 0)
		}

		if robots, _ := robotstxt.FromString(handle.RobotsTXT[r.URL.Host]); robots.TestAgent(r.URL.Path, Config.Name) {
			capturedHTML := HTTPGet(data.URI, 0)
			reader := strings.NewReader(capturedHTML)
			doc, err := goquery.NewDocumentFromReader(reader)
			DeBug("Load HTML", err)
			doc.Find("noscript").Remove() // Remove NoJavascript codes
			doc.Find("script").Remove()   // Remove Javascript codes
			doc.Find("style").Remove()    // Remove CSS codes
			doc.Find("iframe").Remove()   // Remove Iframe codes
			doc.Find("meta").Remove()     // Remove Meta codes
			data.Content = ReplaceSyntaxs(doc.Text(), " ")
		} else {
			forbiddenMsg := "> Forbidden by robots.txt"
			data.Content = forbiddenMsg
			fmt.Println(forbiddenMsg)
		}

		handle.Solr.Update(data)
	})

	collyQueue.AddURL(uri)
	collyQueue.Run(handle.Colly)
}
