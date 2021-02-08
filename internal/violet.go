/*
Package butterfly: The library for butterfly

Copyright(c) 2020 Star Inc. All Rights Reserved.
This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/
package butterfly

// Violet: The default schema for StarStart! Service
// Schema XML: https://bit.ly/2I0QxF8
// Usage: https://bit.ly/32r9Noz

// VioletDataStruct: Data structure for Violet
type VioletDataStruct struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	URI         string `json:"uri"`
	Description string `json:"description"`
	Content     string `json:"content"`
}
