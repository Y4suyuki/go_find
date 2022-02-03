package cmd

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/y4suyuki/go_find/matcher"
)

var rootDir = flag.String("d", ".", "root directory to search files")

func Cmd() {

	flag.Parse()

	root, err := filepath.Abs(*rootDir)

	if err != nil {
		log.Println(err)
	}

	xerr := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if matcher.MatchHiddenFile(path) {
			if d.IsDir() {
				return filepath.SkipDir
			}
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
