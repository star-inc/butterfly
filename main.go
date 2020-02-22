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

type starstartDataStruct struct {
	id          int
	uri         string
	description string
	content     string
}

func main() {
	collyHandle := butterfly.NewCollyClient("")
	collyHandle.Fetch("https://starinc.xyz")
}
