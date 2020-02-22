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

	"github.com/supersonictw/butterfly-solr/solr"
)

// SolrHandle :
type SolrHandle struct {
	Client *solr.SolrInterface
}

type solrConfig struct {
	URI        string `json:"solr_uri"`
	Collection string `json:"collection"`
}

// Config Solr
func readConfig(configPath string) solrConfig {
	var config solrConfig
	jsonFile, err := os.Open(configPath)
	DeBug("Get JSON config", err)
	defer jsonFile.Close()
	srcJSON, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(srcJSON, &config)
	DeBug("Load JSON Initialization", err)
	return config
}

// NewSolrClient :
func NewSolrClient() *SolrHandle {
	handle := new(SolrHandle)
	config := readConfig("config.json")
	client, err := solr.NewSolrInterface(config.URI, config.Collection)
	DeBug("Get Solr Client", err)
	handle.Client = client
	return handle
}

// Query :
func (handle *SolrHandle) Query(request string) *solr.Collection {
	query := solr.NewQuery()
	query.Q(request)
	response := handle.Client.Search(query)
	output, _ := response.Result(nil)
	return output.Results
}

// Update :
func (handle *SolrHandle) Update(data *VioletDataStruct) {
	_, err := handle.Client.UpdateDoc(data)
	DeBug("Solr Update", err)
	handle.Client.Commit()
}
