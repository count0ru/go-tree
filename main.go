package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

var startPath = "./maindir"


func getDirectoryItems(path string) {
	
	listDirectories, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal("Cant read directory", err)
	}
	
	fmt.Println("-")
	
	for idx, item := range listDirectories {
		if item.IsDir() {
		  if idx == len(listDirectories) - 1 {
		 	fmt.Println("└───" + item.Name())		
		  } else {
		  	fmt.Println("├───" + item.Name())		
		  }
		  getDirectoryItems(path + "/" + item.Name())
		}
	}


}

func main() {
	
	fmt.Println("├───" + startPath)
	getDirectoryItems(startPath)

}

