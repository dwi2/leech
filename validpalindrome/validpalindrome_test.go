package validpalindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		input string
		want  bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
	}

	for _, c := range cases {
		got := IsPalindrome(c.input)
		if got != c.want {
			t.Errorf("IsPalindrome(%s) == %v, want %v", c.input, got, c.want)
		}
	}
}
