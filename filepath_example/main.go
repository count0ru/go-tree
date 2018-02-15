package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
)

func printFile(path string, info os.FileInfo, err error) error {
    if err != nil {
        log.Print(err)
        return nil
    }
    fmt.Println(path)
    fmt.Printf("%+v\n",info)
    return nil
}

func main() {
    log.SetFlags(log.Lshortfile)
    dir := os.Args[1]
    err := filepath.Walk(dir, printFile)
    if err != nil {
        log.Fatal(err)
    }
}