/*
Butterfly
===
The web crawler base on Apache Solr for StarStart!.

Copyright(c) 2020 Star Inc. All Rights Reserved.
The software licensed under Mozilla Public License Version 2.0
*/
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/vanng822/go-solr/solr"
)

type solrConfig struct {
	URI        string `json:"solr_uri"`
	Collection string `json:"collection"`
}

func readConfig() solrConfig {
	// Config Solr
	var config solrConfig
	jsonFile, err := os.Open("config.json")
	deBug("Get JSON config", err)
	defer jsonFile.Close()
	srcJSON, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(srcJSON, &config)
	deBug("Load JSON Initialization", err)
	return config
}

func deBug(where string, err error) bool {
	if err != nil {
		fmt.Printf("NepCore Error #%s\nReason:\n%s\n\n", where, err)
		return false
	}
	return true
}

func main() {
	config := readConfig()
	client, err := solr.NewSolrInterface(config.URI, config.Collection)
	deBug("Get Solr Client", err)
	query := solr.NewQuery()
	query.Q("*:*")
	response := client.Search(query)
	output, _ := response.Result(nil)
	fmt.Println(output.Results.Docs)
}
