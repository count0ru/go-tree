package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirWalk(out io.Writer, spath string, withFiles bool, parentPrefix string) error {

	var (
		fullPath        string
		prefix          string
		listDirectories []os.FileInfo
	)

	listAllItems, err := ioutil.ReadDir(spath)
	if err != nil {
		return err
	}

	for _, item := range listAllItems {
		if item.IsDir() {
			listDirectories = append(listDirectories, item)
		} else if withFiles {
			listDirectories = append(listDirectories, item)
		}
	}

	for idx, item := range listDirectories {

		var childParentPrefix string

		fullPath = fmt.Sprintf("%s/%s", spath, item.Name())

		switch idx {
		case len(listDirectories) - 1:
			childParentPrefix = "\t"
			prefix = "└───"
		default:
			childParentPrefix = "│\t"
			prefix = "├───"
		}

		if item.IsDir() {
			fmt.Fprintf(out, "%s%s%s\n", parentPrefix, prefix, item.Name())
			dirWalk(out, fullPath, withFiles, fmt.Sprintf("%s%s", parentPrefix, childParentPrefix))
		} else {
			if withFiles {
				if item.Size() == 0 {
					fmt.Fprintf(out, "%s%s%s (%s)\n", parentPrefix, prefix, item.Name(), "empty")
				} else {
					fmt.Fprintf(out, "%s%s%s (%db)\n", parentPrefix, prefix, item.Name(), item.Size())
				}
			}
		}

	}
	return nil
}

func dirTree(outfile io.Writer, path string, showFiles bool) error {

	err := dirWalk(outfile, path, showFiles, "")
	return err
}
