/*
Package butterfly : The library for butterfly

Copyright(c) 2020 Star Inc. All Rights Reserved.
This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/
package butterfly

import (
	"io/ioutil"
	"net/http"
)

// Violet : The default schema for StarStart! Service

// VioletDataStruct : Data structure for Violet
type VioletDataStruct struct {
	Title       string `json:"title"`
	URI         string `json:"uri"`
	Description string `json:"description"`
	Content     string `json:"content"`
}

func VioletHttpGet(url string) string {
	resp, err := http.Get(url)
	DeBug("GetHTTP", err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	DeBug("ReadHTML", err)
	return string(body)
}
