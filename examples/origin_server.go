package main

import (
	"fmt"
	"github.com/nileshjagnik/metadata"
)

func main() {
	data, err := metadata.GetMetadata("NOVA - 35x15 - Car Of The Future.mpg","tv")
	if err != nil {
		fmt.Println(err)
	} else {
        	fmt.Println(data)
	}
}
