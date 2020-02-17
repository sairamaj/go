package morestrings

import "testing"

func TestReverseRuns(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello World", "dello WorlH"},
		{"Another", "rnotheA"},
		{"", ""},
	}

	for _, c := range cases {
		got := ReverseRuns(c.in)

		if got != c.want {
			t.Errorf("ReverseRuns(%q) == %q want %q", c.in, got, c.want)
		}
	}
}
