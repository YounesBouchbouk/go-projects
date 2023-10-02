package main

func ProductExceptSelf(nums []int) []int {
	size := len(nums)
	result := make([]int, size)
	result[0], result[size-1] = 1, 1

	for i := 1; i < size; i++ {
		result[i] = result[i-1] * nums[i-1]
	}

	rightPro := 1
	for i := size - 2; i >= 0; i-- {
		rightPro *= nums[i+1]
		result[i] *= rightPro
	}

	return result
}