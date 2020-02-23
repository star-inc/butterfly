/*
Package butterfly : The library for butterfly

Copyright(c) 2020 Star Inc. All Rights Reserved.
This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/
package butterfly

import (
	"bytes"
	"fmt"
	"strings"
)

// DeBug : Print errors for debug and report
func DeBug(where string, err error) bool {
	if err != nil {
		fmt.Printf("Butterfly Error #%s\nReason:\n%s\n\n", where, err)
		return false
	}
	return true
}

// ReplaceSyntaxs : Remove space and syntax
func ReplaceSyntaxs(rawString string, filled string) string {
	var output bytes.Buffer
	rawString = strings.ReplaceAll(rawString, " ", "\x1e")
	rawString = strings.ReplaceAll(rawString, "\t", "\x1e")
	rawString = strings.ReplaceAll(rawString, "\n", "\x1e")
	stringSlice := strings.Split(rawString, "\x1e")
	for _, word := range stringSlice {
		if word != "" {
			output.WriteString(word + filled)
		}
	}
	return output.String()
}
