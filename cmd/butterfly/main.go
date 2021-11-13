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
	"sync"

	butterfly "github.com/star-inc/butterfly/internal"
)

var (
	addSiteValue    string
	deleteSiteValue string
)

func init() {
	flag.StringVar(&addSiteValue, "add", "", "Append URL into Site List")
	flag.StringVar(&deleteSiteValue, "del", "", "Remove URL from Site List")
	flag.Usage = usage
}

func usage() {
	fmt.Println("Butterfly - The web crawler base on Apache Solr for StarStart!")
	fmt.Println("(c) 2020 Star Inc. https://starinc.xyz")
	fmt.Print("\nOptional argument:\n\n")
}

func main() {
	if addSiteValue != "" {
		butterfly.AddSite(addSiteValue)
	}

	if deleteSiteValue != "" {
		butterfly.DeleteSite(deleteSiteValue)
	}

	switch flag.Arg(0) {
	case "start":
		taskList := new(sync.WaitGroup)
		taskList.Add(len(butterfly.SiteList))
		for _, siteURI := range butterfly.SiteList {
			go func(siteURI string, taskList *sync.WaitGroup) {
				client := butterfly.NewBody()
				client.Fetch(siteURI)
				taskList.Done()
			}(siteURI, taskList)
		}
		taskList.Wait()
		break
	case "list":
		butterfly.ShowSiteList()
		break
	case "config":
		butterfly.EditConfigWithTextEditor()
		break
	default:
		usage()
	}
}
