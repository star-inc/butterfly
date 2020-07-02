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
	"sync"

	butterfly "./libs"
)

var (
	addSiteValue    string
	deleteSiteValue string
	client          *butterfly.Handles
)

func init() {
	flag.StringVar(&addSiteValue, "add", "", "Append URL into Site List")
	flag.StringVar(&deleteSiteValue, "del", "", "Remove URL from Site List")
	flag.Usage = usage
}

func usage() {
	fmt.Println("\nButterfly")
	fmt.Println("=========")
	fmt.Println("\nThe web crawler base on Apache Solr for StarStart!")
	fmt.Printf("\nUsage: %s start\n\n", os.Args[0])
	fmt.Print("\t\t(c) 2020 Star Inc. https://starinc.xyz\n\n")
	fmt.Print("\nOptional argument:\n\n")
	flag.PrintDefaults()
	fmt.Print("  list\n\tShow Site List\n\n")
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

	if addSiteValue != "" {
		butterfly.AddSite(addSiteValue)
	}

	if deleteSiteValue != "" {
		butterfly.DeleteSite(deleteSiteValue)
	}

	if flag.Arg(0) == "start" {
		taskList := new(sync.WaitGroup)
		taskList.Add(len(butterfly.SiteList))
		for _, siteURI := range butterfly.SiteList {
			go func(siteURI string, taskList *sync.WaitGroup) {
				client = butterfly.NewBody()
				client.Fetch(siteURI)
				taskList.Done()
			}(siteURI, taskList)
		}
		taskList.Wait()
	} else if flag.Arg(0) == "list" {
		butterfly.ShowSiteList()
	} else {
		usage()
	}
}
