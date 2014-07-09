// Copyright 2013, Amahi.  All rights reserved.
// Use of this source code is governed by the
// license that can be found in the LICENSE file.

// Functions for preprocess input

package metadata

import (
	"strings"
)

//Convert non-standard format strings to standard format
func preprocess(MediaName string, Hint string) (title string, mediatype string, err error) {
	//strip off year name

	parts := strings.Split(MediaName, "(")
	result := parts[0]
	for _, s := range parts[1:] {
		yearparts := strings.Split(s, ")")
		l := len(yearparts)
		if l > 1 {
			result += "%20" + yearparts[l-1]
		} else {
			result += "%20" + s
		}
	}
	parts = strings.Split(result, " ")
	result = parts[0]
	if len(parts) > 1 {
		for _, s := range parts[1:] {
			result += "%20" + s
		}
	}
	//strip off the extension
	parts = strings.Split(result, ".")
	if len(parts) > 1 {
		result = parts[0]
		for _, s := range parts[1 : len(parts)-1] {
			result += "%20" + s
		}
	}
	//FIXME - add hint detection
	return result, Hint, nil
}
