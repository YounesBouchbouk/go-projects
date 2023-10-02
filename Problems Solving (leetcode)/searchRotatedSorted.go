package main

func searchRotatedSorted(nums []int, target int) int {

	pivot := getPivot(nums)

	if pivot == -1 {
		return searcheForIt(nums, target)
	}
	found := searcheForIt(nums[pivot:], target)

	if found > -1 {
		return len(nums[:pivot]) + found
	}

	found = searcheForIt(nums[:pivot+1], target)

	if found > -1 {
		return found
	}

	return -1
}

func getPivot(nums []int) int {

	left, right := 0, len(nums)-1

	if nums[left] < nums[right] {
		return -1
	}

	for left < right {

		mid := left + (right-left)/2

		if mid < len(nums)-1 && nums[mid] > nums[mid+1] {
			return mid + 1
		} else if mid > 0 && nums[mid] < nums[mid-1] {
			return mid
		}

		if nums[left] < nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}

	}
	return -1
}

func searcheForIt(arr []int, x int) int {
	left := 0
	right := len(arr) - 1

	for right >= left {
		mid := (left + right) / 2
		if x == arr[mid] {
			return mid
		} else if x < arr[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}