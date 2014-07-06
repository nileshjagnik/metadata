package main

import (
	"fmt"
	"github.com/nileshjagnik/metadata"
)


func main() {
        data,err := metadata.GetMetadata("Walking Dead")
        if err!=nil {
                fmt.Println(err)
        } else {
                if data.Media_type == "movie" {
                        fmt.Println("Title: "+data.Title)
                        fmt.Println("Release Date: "+data.Release_date)
                        
                        // to get image url use base_url + image_size + path_to_image
                        fmt.Println(data.Config.Images.Base_url+"original"+data.Poster_path)
                } else {
                        fmt.Println("Name: "+data.Name)
                        fmt.Println("Last Air Date: "+data.Last_air_date)
                        fmt.Println(data.Config.Images.Base_url+"original"+data.Poster_path)
                }
        }
}
