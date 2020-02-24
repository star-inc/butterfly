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
		fmt.Println("\n\tButterfly")
		fmt.Println("\t ===== ")
		fmt.Println("\n\tThe web crawler base on Apache Solr for StarStart!")
		fmt.Printf("\n\tUsage: %s <URI>\n\n", os.Args[0])
		fmt.Print("\t\t(c) 2020 Star Inc. https://starinc.xyz\n\n")
	}
}
