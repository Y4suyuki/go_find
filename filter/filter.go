package filter

import "github.com/y4suyuki/go_find/matcher"

func Filter(exp string, strs []string) []string {
	return []string{"foo", "bar"}
}

func FilterHiddenFile(strs []string) []string {
	var res []string
	for _, p := range strs {
		if !matcher.MatchHiddenFile(p) {
			res = append(res, p)
		}
	}

	return res
}
