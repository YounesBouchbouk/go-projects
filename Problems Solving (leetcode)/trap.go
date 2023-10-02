package main

func trap(height []int) int {
	count := 0
	left := 0
	right := len(height) - 1
	maxLeft := height[left]
	maxRight := height[right]

	if len(height) == 0 {
		return 0
	}

	for left < right {
		if maxLeft < maxRight {
			left++
			maxLeft = maxx(maxLeft, height[left])
			count += maxLeft - height[left]

		} else {
			right--
			maxRight = maxx(maxRight, height[right])
			count += maxRight - height[right]

		}
	}

	return count
}

func maxx(a, b int) int {
	if a > b {
		return a
	}
	return b
}
