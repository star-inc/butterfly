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

const (
	configFileName   string = "/config.json"
	siteListFileName string = "/sites.json"
)

// ConfigPath :
var ConfigPath string

// SiteList :
var SiteList []string

// Config : Global Settings for butterfly from config.json
var Config configStruct

// Initiate : Load configure file to Config
func Initiate() {
	readSiteList()
	readConfig()
}

func readSiteList() {
	jsonFile, err := os.Open(ConfigPath + siteListFileName)
	DeBug("Get JSON config", err)
	defer jsonFile.Close()
	srcJSON, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(srcJSON, &SiteList)
	DeBug("Load JSON Initialization", err)
}

// WriteSiteList :
func WriteSiteList() {
	file, _ := json.Marshal(SiteList)
	_ = ioutil.WriteFile(ConfigPath+siteListFileName, file, 0644)
}

func readConfig() {
	jsonFile, err := os.Open(ConfigPath + configFileName)
	DeBug("Get JSON config", err)
	defer jsonFile.Close()
	srcJSON, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(srcJSON, &Config)
	DeBug("Load JSON Initialization", err)
}
