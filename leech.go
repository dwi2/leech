package main

import (
	"fmt"
	"leech/mergesortedarray"
)

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	actual := mergesortedarray.Merge(nums1, 3, nums2, 3)
	// expect []int{1, 2, 2, 3, 5, 6}
	fmt.Printf("%v\n", actual)
}
