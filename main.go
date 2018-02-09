package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	
	listDirectories, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal("Cant read directory", err)
	}

	for _, item := range listDirectories {
		if item.IsDir() {
		  fmt.Println(item.Name())		
		}
	}
}

