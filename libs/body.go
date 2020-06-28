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

// NewBody : Create body for butterfly
func NewBody() *Handles {
	Initiate()
	handle := new(Handles)
	handle.Colly = colly.NewCollector(
		colly.UserAgent(Config.UserAgent),
	)
	handle.Solr = NewSolrClient()
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

func (handle *Handles) collect(uri string) *VioletDataStruct {
	data := new(VioletDataStruct)

	data.URI = uri
	queryURL, _ := url.Parse(uri)

	if queryURL.Scheme == "" {
		data.URI = fmt.Sprintf("http:%s", uri)
		queryURL, _ = url.Parse(data.URI)
	}

	fmt.Println("Collecting", data.URI)

	signature := md5.Sum([]byte(data.URI))
	data.ID = fmt.Sprintf("%x", signature)

	if _, exists := handle.RobotsTXT[queryURL.Host]; !exists {
		handle.RobotsTXT[queryURL.Host] = HTTPGet(fmt.Sprintf("%s://%s/robots.txt", queryURL.Scheme, queryURL.Host), 0)
	}

	if robots, _ := robotstxt.FromString(handle.RobotsTXT[queryURL.Host]); robots.TestAgent(queryURL.Path, Config.Name) {
		capturedHTML := HTTPGet(data.URI, 0)

		reader := strings.NewReader(capturedHTML)
		doc, err := goquery.NewDocumentFromReader(reader)
		DeBug("Load HTML", err)

		data.Title = doc.Find("title").Text()
		data.Description, _ = doc.Find("meta[name=description]").Attr("content")

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

	return data
}

// Fetch : Capture web pages on Internet and submit to Solr
func (handle *Handles) Fetch(uri string) {
	queryURL, _ := url.Parse(uri)
	handle.setStorage(queryURL.Host)
	defer handle.CollyStorage.Close()

	if Config.ForcusOnURI {
		handle.Colly.AllowedDomains = []string{queryURL.Host}
	}

	var collyQueue *queue.Queue
	collyQueue, _ = queue.New(Config.Colly.Threads, handle.CollyStorage)

	handle.Colly.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
		data := handle.collect(r.URL.String())
		handle.Solr.Update(data)
	})

	handle.Colly.OnHTML("a[href]", func(e *colly.HTMLElement) {
		collyQueue.AddURL(e.Request.AbsoluteURL(e.Attr("href")))
	})

	collyQueue.AddURL(uri)
	collyQueue.Run(handle.Colly)
}
