package matcher

import (
	"fmt"
	"regexp"
)

func MatchHiddenFile(path string) bool {
	m, err := regexp.MatchString(`^(.*/)?\.[^./].*`, path)
	if err != nil {
		fmt.Errorf("Error for path: %s", path)
	}

	return m
}
