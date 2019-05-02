package main

import (
	"fmt"
	"leech/mergesortedarray"
)

func main() {
	nums1 := []int{1, 2, 4, 5, 6, 0}
	nums2 := []int{3}
	actual := mergesortedarray.Merge(nums1, 5, nums2, 1)
	// expect []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("%v\n", actual)
}
