/*
Package butterfly : The library for butterfly

Copyright(c) 2020 Star Inc. All Rights Reserved.
This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/
package butterfly

import (
	"github.com/gocolly/colly"
)

// CollyHandle :
type CollyHandle struct {
	Client    *colly.Collector
	UserAgent string
}

// NewCollyClient :
func NewCollyClient(userAgent string) {
	handle := new(CollyHandle)
	handle.setUserAgent(userAgent)
	client := colly.NewCollector(
		colly.UserAgent(handle.UserAgent),
	)
	handle.Client = client
}

func (handle *CollyHandle) setUserAgent(userAgent string) {
	if userAgent == "" {
		handle.UserAgent = "Mozilla/5.0 (compatible; Star Butterfly/1.0; +https://github.com/star-inc/butterfly)"
	} else {
		handle.UserAgent = userAgent
	}
}
