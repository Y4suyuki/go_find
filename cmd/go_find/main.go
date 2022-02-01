package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/y4suyuki/go_find/matcher"
)

func Cmd() {
	root, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	var files []string

	xerr := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if matcher.MatchHiddenFile(path) {
			return nil
		}
		out_path, err := filepath.Rel(root, path)
		if err != nil {
			fmt.Errorf("Oh no!, %e", err)
		}
		files = append(files, out_path)
		return nil
	})

	if xerr != nil {
		panic(err)
	}

	for _, file := range files {
		fmt.Println(file)
	}

	fmt.Println(root)
}
