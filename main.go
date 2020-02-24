/*
Butterfly
===
The web crawler base on Apache Solr for StarStart!.

Copyright(c) 2020 Star Inc. All Rights Reserved.
The software licensed under Mozilla Public License Version 2.0
*/
package main

import (
	"fmt"
	"os"

	"./butterfly"
)

func main() {
	butterfly.ReadConfig("config.json")
	solrHandle := butterfly.NewSolrClient()
	collyHandle := butterfly.NewCollyClient(solrHandle)
	if len(os.Args) == 2 {
		collyHandle.Fetch(os.Args[1])
	} else {
		fmt.Printf("Usage: %s <URI>\n", os.Args[0])
	}
}
