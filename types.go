// Copyright 2013, Amahi.  All rights reserved.
// Use of this source code is governed by the
// license that can be found in the LICENSE file.

// Data structures and types for the metadata library

package metadata

const API_KEY string = "dc10d9b00f8a4a777539655342cbb647"

//response of search/multi
type TmdbResponse struct {
	Page          int
	Results       []TmdbResult
	Total_pages   int
	Total_results int
}

type TmdbResult struct {
	Adult          bool
	Name           string
	Backdrop_path  string
	Id             int
	Original_name  string
	Original_title string
	First_air_date string
	Release_date   string
	Poster_path    string
	Title          string
	Media_type     string
	Profile_path   string
}

//response of config
type TmdbConfig struct {
	Images      ImageConfig
	Change_keys []string
}

type ImageConfig struct {
	Base_url        string
	Secure_base_url string
	
	//possible sizes for images
	Backdrop_sizes  []string
	Logo_sizes      []string
	Poster_sizes    []string
	Profile_sizes   []string
	Still_sizes     []string
}

//final metadata to be returned
type Metadata struct {
	//common fields
	Id             int
	Backdrop_path  string
	Poster_path    string
	Credits        TmdbCredits
	Media_type     string
	Config         TmdbConfig
	
	//movie specific fields
	Imdb_id        string
	Overview       string
	Title          string
	Release_date   string
	
	//tv specific fields
	Name           string
	First_air_date string
	Last_air_date  string
}

type TmdbCredits struct {
	Id   int
	Cast []TmdbCast
	Crew []TmdbCrew
}

type TmdbCast struct {
	Character    string
	Name         string
	Profile_path string
}

type TmdbCrew struct {
	Department   string
	Name         string
	Job          string
	Profile_path string
}
