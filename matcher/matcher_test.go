package matcher

import "testing"

func TestMatcher(t *testing.T) {

	matcherTests := []struct {
		p    string
		want bool
	}{
		{"foo", false},
		{"foo.txt", false},
		{"./foo", false},
		{"../hello/world", false},
		// matched as hidden files
		{".foo", true},
		{"/.foo", true},
		{"$HOME/.foo", true},
		{"/foo/bar/.foo/baz", true},
	}

	for _, tt := range matcherTests {
		got := MatchHiddenFile(tt.p)
		if got != tt.want {
			t.Errorf("%s expects result %t but %t", tt.p, tt.want, got)
		}
	}

}
