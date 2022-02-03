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

	xerr := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if matcher.MatchHiddenFile(path) {
			return nil
		}
		out_path, err := filepath.Rel(root, path)
		if err != nil {
			fmt.Errorf("Oh no!, %e", err)
		}
		fmt.Println(out_path)
		return nil
	})

	if xerr != nil {
		panic(err)
	}

	fmt.Println(root)
}
