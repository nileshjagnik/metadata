package main

import (
	"fmt"
	"github.com/nileshjagnik/metadata"
)


func main() {
        data,err := metadata.GetInfo("Fight%20Club")
        if err!=nil {
                fmt.Println(err)
        } else {
                fmt.Println(data)
        }
}
