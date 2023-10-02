package main

func searchMatrix(matrix [][]int, target int) bool {

	for i := 0; i < len(matrix); i++ {
		if search(matrix[i], target) != -1 {
			return true
		}
	}

	return false
}

func search(arr []int, x int) int {
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
