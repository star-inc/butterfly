// Butterfly - The web crawler base on Apache Solr for StarStart!
// Copyright(c) 2020 Star Inc. All Rights Reserved.
// The software licensed under Mozilla Public License Version 2.0

package spec

// Violet: The default schema for StarStart! Service
// Schema XML: https://bit.ly/2I0QxF8
// Usage: https://bit.ly/32r9Noz

// Violet Data structure
type Violet struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	URI         string `json:"uri"`
	Description string `json:"description"`
	Content     string `json:"content"`
}
