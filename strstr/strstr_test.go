package strstr

import "testing"

func TestStrStr(t *testing.T) {
	cases := []struct {
		haystack, needle string
		want             int
	}{
		{"Hello", "ll", 2},
		{"World", "Z", -1},
		{"Empty needle", "", 0},
		{"aaa", "aaaa", -1},
		{"aaaaa", "bba", -1},
		{"mississippi", "issip", 4},
		{"mississippi", "issipi", -1},
	}
	for _, c := range cases {
		got := StrStr(c.haystack, c.needle)
		if got != c.want {
			t.Errorf("strStr(%q, %q) == %d, want %d", c.haystack, c.needle, got, c.want)
		}
	}
}
