/*
Butterfly
===
The web crawler base on Apache Solr for StarStart!.

Copyright(c) 2020 Star Inc. All Rights Reserved.
The software licensed under Mozilla Public License Version 2.0
*/
package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"

	butterfly "./libs"
)

var (
	addSiteValue    string
	deleteSiteValue string
	client          *butterfly.Handles
)

func init() {
	flag.StringVar(&addSiteValue, "add", "", "the action code")
	flag.StringVar(&deleteSiteValue, "del", "", "the action code")
	flag.Usage = usage
}

func usage() {
	fmt.Println("\n\tButterfly")
	fmt.Println("\t ===== ")
	fmt.Println("\n\tThe web crawler base on Apache Solr for StarStart!")
	fmt.Printf("\n\tUsage: %s <URI>\n\n", os.Args[0])
	fmt.Print("\t\t(c) 2020 Star Inc. https://starinc.xyz\n\n")
	flag.PrintDefaults()
}

func addSite(domain string) {
	fmt.Println(domain)
}

func deleteSite(domain string) {
	fmt.Println(domain)
}

func fly() {
	for i := range butterfly.SiteList {
		go client.Fetch(butterfly.SiteList[i].Domain + butterfly.SiteList[i].StartPath)
	}
}

func main() {
	var configPathRoot string
	flag.Parse()
	osUser, err := user.Current()
	butterfly.DeBug("Get OS User", err)
	if osUser.Username == "root" {
		configPathRoot = "/etc"
	} else {
		configPathRoot = osUser.HomeDir
	}
	configPath := fmt.Sprintf("%s/.config/butterfly", configPathRoot)
	butterfly.ReadConfig(configPath)
	butterfly.ReadSiteList(configPath)
	client = butterfly.NewBody()

	if addSiteValue != "" {
		addSite(addSiteValue)
	}

	if deleteSiteValue != "" {
		deleteSite(deleteSiteValue)
	}

	if flag.Arg(0) == "start" {
		fly()
	} else {
		usage()
	}
}
