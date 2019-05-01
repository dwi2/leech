package mergesortedarray

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	cases := []struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}{
		{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
	}
	for _, c := range cases {
		got := Merge(c.nums1, c.m, c.nums2, c.n)
		if reflect.DeepEqual(got, c.want) != true {
			t.Errorf("Merge() == %v, want %v", got, c.want)
		}
	}

}
