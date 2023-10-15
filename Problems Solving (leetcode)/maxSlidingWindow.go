// maxSlidingWindow problem
package main

func FindMaxFromlist(nums []int) int {
	max := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}

	return max
}

func MaxSlidingWindow(nums []int, k int) []int {

	l, r := 0, k-1
	var result []int

	for r < len(nums) {
		result = append(result, FindMaxFromlist(nums[l:r+1]))
		l++
		r++
	}

	return result
}
