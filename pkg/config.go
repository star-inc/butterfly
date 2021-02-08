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
	"fmt"
	"io/ioutil"
	"os"
)

type configStruct struct {
	Name             string      `json:"name"`
	FocusOnURIDomain bool        `json:"focus-on-uri-domain"`
	UserAgent        string      `json:"user-agent"`
	Solr             solrConfig  `json:"solr"`
	Colly            collyConfig `json:"colly"`
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

var configPath string

// SiteList : The website list that will be fetched by the butterfly
var SiteList []string

// Config : Global Settings for butterfly from config.json
var Config configStruct

// Initiate : Load configure file to Config
func Initiate(path string) {
	configPath = path
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		err = os.MkdirAll(configPath, 0755)
		DeBug("Config Initiate create directory", err)
	}
	readConfig()
	readSiteList()
}

func readConfig() {
	jsonFile, err := os.Open(configPath + configFileName)
	DeBug("Get JSON config", err)
	defer jsonFile.Close()
	srcJSON, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(srcJSON, &Config)
	DeBug("Load JSON Initialization", err)
}

func readSiteList() {
	jsonFile, err := os.Open(configPath + siteListFileName)
	DeBug("Get JSON config", err)
	defer jsonFile.Close()
	srcJSON, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(srcJSON, &SiteList)
	DeBug("Load JSON Initialization", err)
}

// ShowSiteList : To show the website list that will be fetched
func ShowSiteList() {
	for _, siteURI := range SiteList {
		fmt.Println(siteURI)
	}
	os.Exit(0)
}

// AddSite : Add a website for fetching
func AddSite(siteURI string) {
	URI, _ := NormalizeURI(siteURI)
	_, exists := FindInSlice(SiteList, URI)
	if !exists {
		SiteList = append(SiteList, URI)
		fmt.Println("Successful.")
	} else {
		fmt.Printf("%s already existed in SiteList.\n", URI)
	}
	WriteSiteList()
	os.Exit(0)
}

// DeleteSite : Remove a website from the website list.
func DeleteSite(siteURI string) {
	URI, _ := NormalizeURI(siteURI)
	var newSiteList []string
	index, exists := FindInSlice(SiteList, URI)
	if exists {
		for i, item := range SiteList {
			if i != index {
				newSiteList = append(newSiteList, item)
			}
		}
		SiteList = newSiteList
		fmt.Println("Successful.")
	} else {
		fmt.Printf("%s not exists in SiteList.\n", URI)
	}
	WriteSiteList()
	os.Exit(0)
}

// EditConfigWithTextEditor : Call the text editor to modify the config file.
func EditConfigWithTextEditor() {
	CallTextEditor(configPath + configFileName)
}

// WriteSiteList : Make the butterfly write the website list after it was modified
func WriteSiteList() {
	file, _ := json.Marshal(SiteList)
	_ = ioutil.WriteFile(configPath+siteListFileName, file, 0644)
}
