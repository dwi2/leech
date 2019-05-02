package mergesortedarray

func Merge(nums1 []int, m int, nums2 []int, n int) []int {
	for i := 0; i < len(nums2); i++ {
		nums1[m+i] = nums2[i]
	}
	quicksort(nums1, 0, len(nums1)-1)
	return nums1
}

func quicksort(nums []int, low int, hi int) {
	if low < hi {
		pivot := partition(nums, low, hi)
		quicksort(nums, low, pivot-1)
		quicksort(nums, pivot+1, hi)
	}
	return
}

func partition(nums []int, low int, hi int) int {
	pivot := nums[hi]
	i := low
	for j := low; j < hi; j++ {
		if nums[j] < pivot {
			swap(nums, i, j)
			i++
		}
	}
	swap(nums, i, hi)
	return i
}

func swap(nums []int, i, j int) {
	if i < 0 || j < 0 || i >= len(nums) || j >= len(nums) {
		return
	}
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}
