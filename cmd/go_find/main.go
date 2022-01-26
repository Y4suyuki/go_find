package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	var files []string

	xerr := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})

	if xerr != nil {
		panic(err)
	}

	for _, file := range files {
		fmt.Println(file)
	}

	fmt.Println(path)
}
