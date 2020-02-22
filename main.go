/*
Butterfly
===
The web crawler base on Apache Solr for StarStart!.

Copyright(c) 2020 Star Inc. All Rights Reserved.
The software licensed under Mozilla Public License Version 2.0
*/
package main

import (
	"./butterfly"
)

func main() {
	collyHandle := butterfly.NewCollyClient("")
	solrHandle := butterfly.NewSolrClient()
	collyHandle.Fetch("https://starinc.xyz", solrHandle)
}
