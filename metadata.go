// Copyright 2013, Amahi.  All rights reserved.
// Use of this source code is governed by the
// license that can be found in the LICENSE file.

// Functions for requesting Metadata

package metadata

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

var config = TmdbConfig{}

//search for TV, persons and Movies with a given name
func SearchMulti(MediaName string) (TmdbResponse, error) {
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
		return TmdbResponse{}, err
	}
	return resp, nil
}

//search for Movies with a given name
func SearchMovie(MediaName string) (TmdbResponse, error) {
	res, err := http.Get("http://api.themoviedb.org/3/search/movie?api_key=" + API_KEY + "&query=" + MediaName)
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
		return TmdbResponse{}, err
	}
	return resp, nil
}

//search for Tv Shows with a given name
func SearchTv(MediaName string) (TmdbResponse, error) {
	res, err := http.Get("http://api.themoviedb.org/3/search/tv?api_key=" + API_KEY + "&query=" + MediaName)
	var resp TmdbResponse
	if err != nil {
		return resp, err
	}
	if res.StatusCode != 200 {
		return resp, errors.New("Status Code 200 not recieved from TMDb")
	}
	body, err := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return TmdbResponse{}, err
	}
	return resp, nil
}

//get configurations
func GetConfig() (TmdbConfig, error) {
        if config.Change_keys == nil {
                res, err := http.Get("http://api.themoviedb.org/3/configuration?api_key=" + API_KEY)
	        var conf TmdbConfig
	        if err != nil {
		        return conf, err
	        }
	        if res.StatusCode != 200 {
		        return conf, errors.New("Status Code 200 not recieved from TMDb")
	        }
	        body, err := ioutil.ReadAll(res.Body)
	        err = json.Unmarshal(body, &conf)
	        if err != nil {
		        return TmdbConfig{}, err
	        }
	        config = conf
	        return conf, nil
        } else {
                return config,nil
        }
}

//get basic information for movie
func GetMovieDetails(MediaId string) (Metadata, error) {
        res, err := http.Get("http://api.themoviedb.org/3/movie/"+MediaId+"?api_key=" + API_KEY)
	var met Metadata
	if err != nil {
		return met, err
	}
	if res.StatusCode != 200 {
		return met, errors.New("Status Code 200 not recieved from TMDb")
	}
	body, err := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &met)
	if err != nil {
		return Metadata{}, err
	}
	return met, nil
}

//get credits for movie
func GetMovieCredits(MediaId string) (TmdbCredits, error) {
        res, err := http.Get("http://api.themoviedb.org/3/movie/"+MediaId+"/credits?api_key=" + API_KEY)
	var cred TmdbCredits
	if err != nil {
		return cred, err
	}
	if res.StatusCode != 200 {
		return cred, errors.New("Status Code 200 not recieved from TMDb")
	}
	body, err := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &cred)
	if err != nil {
		return TmdbCredits{}, err
	}
	return cred, nil
}

//get basic information for Tv
func GetTvDetails(MediaId string) (Metadata, error) {
        res, err := http.Get("http://api.themoviedb.org/3/tv/"+MediaId+"?api_key=" + API_KEY)
	var met Metadata
	if err != nil {
		return met, err
	}
	if res.StatusCode != 200 {
		return met, errors.New("Status Code 200 not recieved from TMDb")
	}
	body, err := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &met)
	if err != nil {
		return Metadata{}, err
	}
	return met, nil
}

//get credits for Tv
func GetTvCredits(MediaId string) (TmdbCredits, error) {
        res, err := http.Get("http://api.themoviedb.org/3/tv/"+MediaId+"/credits?api_key=" + API_KEY)
	var cred TmdbCredits
	if err != nil {
		return cred, err
	}
	if res.StatusCode != 200 {
		return cred, errors.New("Status Code 200 not recieved from TMDb")
	}
	body, err := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &cred)
	if err != nil {
		return TmdbCredits{}, err
	}
	return cred, nil
}

//get all metadata for any kind of media
func GetMetadata(MediaName string) (Metadata, error) {
        var met Metadata
        results,err := SearchMulti(preprocess(MediaName))
        if err != nil {
                return met,err
        }
        if results.Total_results == 0 {
                return met,errors.New("No results found at TMDb")
        }
        if results.Results[0].Media_type == "person" {
                return met,errors.New("Metadata for persons not supported")
        } else if results.Results[0].Media_type == "tv" {
                met,err = GetTvDetails(strconv.Itoa(results.Results[0].Id))
                if err != nil {
                        return met,err
                }
                met.Credits,err = GetTvCredits(strconv.Itoa(results.Results[0].Id))
                if err != nil {
                        return met,err
                }
                met.Config,err = GetConfig()
                if err != nil {
                        return met,err
                }
                met.Media_type = "tv"
                met.Id = results.Results[0].Id
                return met,nil
        } else {
                met,err = GetMovieDetails(strconv.Itoa(results.Results[0].Id))
                if err != nil {
                        return met,err
                }
                met.Credits,err = GetMovieCredits(strconv.Itoa(results.Results[0].Id))
                if err != nil {
                        return met,err
                }
                met.Config,err = GetConfig()
                if err != nil {
                        return met,err
                }
                met.Media_type = "movie"
                met.Id = results.Results[0].Id
                return met,nil
        }
}

//preprocess string
func preprocess(MediaName string) string {
        parts := strings.Split(MediaName, " ")
        result := parts[0]
        if len(parts)>1 {
                for _,s:= range parts[1:] {
                        result += "%20"+s
                }
        }
        return result
}
