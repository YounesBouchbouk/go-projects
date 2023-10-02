package main

func findMin(nums []int) int {
	left, right := 0, len(nums)-1

	if len(nums) == 1 {
		return nums[0]
	}
	if nums[left] < nums[right] {
		return nums[left]
	}

	for left < right {

		mid := left + (right-left)/2

		if mid-1 >= 0 {
			if nums[mid] < nums[mid-1] {
				return nums[mid]
			}
		}

		if nums[mid] < nums[right] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return nums[left]
}
