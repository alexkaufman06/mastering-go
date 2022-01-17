package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// Example command: go run which.go fmt test log

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide an argument!")
		return
	}
	files := arguments[1:]
	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)

	for _, file := range files {
		for _, directory := range pathSplit {
			fullPath := filepath.Join(directory, file)
			fileInfo, err := os.Stat(fullPath) // Does it exist?
			if err == nil {
				mode := fileInfo.Mode()
				if mode.IsRegular() {
					if mode&0111 != 0 { // Is executable?
						fmt.Println(fullPath)
					}
				}
			}
		}
	}
}