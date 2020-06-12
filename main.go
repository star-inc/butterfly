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
	"os/user"

	butterfly "./libs"
)

func main() {
	var configPathRoot string
	osUser, err := user.Current()
	butterfly.DeBug("Get OS User", err)
	if osUser.Username == "root" {
		configPathRoot = "/etc"
	} else {
		configPathRoot = osUser.HomeDir
	}
	configPath := fmt.Sprintf("%s/.config/butterfly/config.json", configPathRoot)
	butterfly.ReadConfig(configPath)
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
