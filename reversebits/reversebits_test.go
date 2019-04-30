package reversebits

import "testing"

func TestReverseBits(t *testing.T) {
	cases := []struct {
		num  uint32
		want uint32
	}{
		{43261596, 964176192},
		{4294967293, 3221225471},
	}
	for _, c := range cases {
		got := ReverseBits(c.num)
		if got != c.want {
			t.Errorf("ReverseBits(%d) == %d, want %d", c.num, got, c.want)
		}
	}

}
