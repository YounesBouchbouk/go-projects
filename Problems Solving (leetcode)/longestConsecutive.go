package main

func longestConsecutive(nums []int) int {

	longest := 0
	numSet := make(map[int]bool)

	for _, val := range nums {
		numSet[val] = true
	}

	for i := 0; i < len(nums); i++ {
		lenght := 0
		if !numSet[nums[i]-1] {
			for numSet[nums[i]+lenght] {
				lenght++
			}

			if lenght > longest {
				longest = lenght
			}
		}
	}

	return longest
}
