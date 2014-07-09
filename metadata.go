// Copyright 2013, Amahi.  All rights reserved.
// Use of this source code is governed by the
// license that can be found in the LICENSE file.

// Functions for requesting Metadata

package metadata

import ()

//get metadata for Tv shows and movies from a given medianame and hint as to whether the media is "tv" or "movie"
func GetMetadata(MediaName string, Hint string) (json string, err error) {
	var met string
	processed_string, mediatype, err := preprocess(MediaName, Hint)
	if err != nil {
		return met, err
	}
	met, err = cache_lookup(processed_string, mediatype)
	if err == nil {
		return met, err
	} else {
		if mediatype == "tv" {
			met, err := getTvData(processed_string)
			if err != nil {
				return "", err
			}
			return met, err
		} else if mediatype == "movie" {
			met, err := getMovieData(processed_string)
			if err != nil {
				return "", err
			}
			return met, err
		}
	}
	return met, err
}
