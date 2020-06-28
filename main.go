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
	var newSite butterfly.SiteListItem
	newSite.Domain = domain
	_, exists := butterfly.FindInSlice(butterfly.SiteList, newSite)
	if !exists {
		butterfly.SiteList = append(butterfly.SiteList, newSite)
	} else {
		fmt.Printf("%s already existed in SiteList.\n", domain)
	}
	butterfly.WriteSiteList()
	os.Exit(0)
}

func deleteSite(domain string) {
	var newSiteList []butterfly.SiteListItem
	var newSite butterfly.SiteListItem
	newSite.Domain = domain
	index, exists := butterfly.FindInSlice(butterfly.SiteList, newSite)
	if exists {
		for i, item := range butterfly.SiteList {
			if i != index {
				newSiteList = append(newSiteList, item)
			}
		}
		butterfly.SiteList = newSiteList
	} else {
		fmt.Printf("%s not exists in SiteList.\n", domain)
	}
	butterfly.WriteSiteList()
	os.Exit(0)
}

func fly() {
	for _, item := range butterfly.SiteList {
		go client.Fetch(item.Domain + item.StartPath)
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
	butterfly.ConfigPath = fmt.Sprintf("%s/.config/butterfly", configPathRoot)
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
