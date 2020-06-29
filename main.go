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

func showSiteList() {
	for _, siteURI := range butterfly.SiteList {
		fmt.Println(siteURI)
	}
	os.Exit(0)
}

func addSite(siteURI string) {
	_, exists := butterfly.FindInSlice(butterfly.SiteList, siteURI)
	if !exists {
		butterfly.SiteList = append(butterfly.SiteList, siteURI)
	} else {
		fmt.Printf("%s already existed in SiteList.\n", siteURI)
	}
	butterfly.WriteSiteList()
	os.Exit(0)
}

func deleteSite(siteURI string) {
	var newSiteList []string
	index, exists := butterfly.FindInSlice(butterfly.SiteList, siteURI)
	if exists {
		for i, item := range butterfly.SiteList {
			if i != index {
				newSiteList = append(newSiteList, item)
			}
		}
		butterfly.SiteList = newSiteList
	} else {
		fmt.Printf("%s not exists in SiteList.\n", siteURI)
	}
	butterfly.WriteSiteList()
	os.Exit(0)
}

func getConfigPath() string {
	var configPathRoot string
	flag.Parse()
	osUser, err := user.Current()
	butterfly.DeBug("Get OS User", err)
	if osUser.Username == "root" {
		configPathRoot = "/etc"
	} else {
		configPathRoot = osUser.HomeDir
	}
	return fmt.Sprintf("%s/.config/butterfly", configPathRoot)
}

func main() {
	butterfly.ConfigPath = getConfigPath()
	client = butterfly.NewBody()

	if addSiteValue != "" {
		addSite(addSiteValue)
	}

	if deleteSiteValue != "" {
		deleteSite(deleteSiteValue)
	}

	if flag.Arg(0) == "start" {
		for _, siteURI := range butterfly.SiteList {
			client.Fetch(siteURI)
		}
	} else if flag.Arg(0) == "list" {
		showSiteList()
	} else {
		usage()
	}
}
