package main

import "sort"

func threeSum(nums []int) [][]int {
	result := [][]int{}

	sort.Ints(nums)

	for i, a := range nums {
		if i > 0 && a == nums[i-1] {
			continue
		}

		l := i + 1
		r := len(nums) - 1

		for l < r {
			sumS := a + nums[l] + nums[r]

			if sumS > 0 {
				r--
			} else if sumS < 0 {
				l++
			} else {
				rs := []int{a, nums[l], nums[r]}
				result = append(result, rs)
				l++
				for nums[l] == nums[l-1] && l < r {
					l++
				}
			}
		}

	}

	return result

}
