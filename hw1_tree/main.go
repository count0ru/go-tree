package main

import (
	"fmt"
//	"io"
	"io/ioutil"
	"os"
//	"path/filepath"
	"strings"
)

func main() {
	//out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirWalk(spath string, withFiles bool ) error {

	listDirectories, err := ioutil.ReadDir(spath)
	if err != nil {
		return err
	}
	
	var fullPath string
	var prefix string
	var level int
	
	for idx, item := range listDirectories {
		
		fullPath = spath + "/" + item.Name()
		level = strings.Count(fullPath,"/")-2
		prefix =  strings.Repeat("|\t", level)

		switch idx {
			case len(listDirectories)-1:
				prefix = prefix + "└───"
			default:
				prefix = prefix + "├───" 
		}

		if item.IsDir() {
			fmt.Println(prefix ,item.Name())
			dirWalk(fullPath, withFiles)
 	 	} else if withFiles {
			fmt.Println(prefix, item.Name())
		}
  	}

}


func dirTree(path string, showFiles bool) error {

	err := dirWalk(path, showFiles)	
	return err
}
