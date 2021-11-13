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
}

func main() {
	switch flag.Arg(0) {
	case "start":
	case "list":
	case "config":
	default:
		usage()
	}
}
