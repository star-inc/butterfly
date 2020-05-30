/*
Package butterfly : The library for butterfly

Copyright(c) 2020 Star Inc. All Rights Reserved.
This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/
package butterfly

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type siteList struct {
	Domain string `json:"domain"`
	StartPath string `json:"start-path"`
}

type configStruct struct {
	Name        string      `json:"name"`
	ForcusOnURI bool        `json:"forcus-on-uri"`
	UserAgent   string      `json:"user-agent"`
	Solr        solrConfig  `json:"solr"`
	Colly       collyConfig `json:"colly"`
}

type solrConfig struct {
	URI        string `json:"uri"`
	Collection string `json:"collection"`
}

type collyConfig struct {
	Threads    int    `json:"threads"`
	SqlitePath string `json:"sqlite_path"`
}

// SiteList : 
var SiteList siteList

// Config : Global Settings for butterfly from config.json
var Config configStruct

// ReadSiteList : 
func ReadSiteList(configPath string) {
	jsonFile, err := os.Open(configPath)
	DeBug("Get JSON config", err)
	defer jsonFile.Close()
	srcJSON, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(srcJSON, &SiteList)
	DeBug("Load JSON Initialization", err)
}

// ReadConfig : Load configure file to Config
func ReadConfig(configPath string) {
	jsonFile, err := os.Open(configPath)
	DeBug("Get JSON config", err)
	defer jsonFile.Close()
	srcJSON, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(srcJSON, &Config)
	DeBug("Load JSON Initialization", err)
}
