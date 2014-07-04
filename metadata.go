// Copyright 2013, Amahi.  All rights reserved.
// Use of this source code is governed by the
// license that can be found in the LICENSE file.

// Functions for requesting Metadata

package metadata

import (
        "encoding/json"
        "errors"
        "fmt"
	"io/ioutil"
	"net/http"
)

const API_KEY string = "dc10d9b00f8a4a777539655342cbb647"

func GetInfo(MediaName string) (TmdbResponse, error) {
	res, err := http.Get("http://api.themoviedb.org/3/search/multi?api_key=" + API_KEY + "&query=" + MediaName)
	var resp TmdbResponse
	if err != nil {
		return resp, err
	}
	if res.StatusCode != 200 {
	        return resp, errors.New("Status Code 200 not recieved from TMDB") 
	}
	body, err := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &resp)
	if err != nil {
		fmt.Println("error:", err)
	}
	return resp, nil
}
