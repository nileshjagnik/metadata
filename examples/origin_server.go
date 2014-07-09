package main

import (
	"fmt"
	"github.com/nileshjagnik/metadata"
)

func main() {
	data, err := metadata.GetMetadata("Robin Williams Weapons of Self Destruction (2009).avi","movie")
	if err != nil {
		fmt.Println(err)
	} else {
        	fmt.Println(data)
	}
}
