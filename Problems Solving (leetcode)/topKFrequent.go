package main

import (
	"sort"
)

// fmt.Println("Map-2: ", topKFrequent([]int{2 , 2, 3 , 3 , 3 , 4 , 6}, 1))

func topKFrequent(nums []int, k int) []int {
	mymap := make(map[int]int)

	for _, i := range nums {
		mymap[i] = mymap[i] + 1
	}

	if len(nums) == 1 {
		return nums
	}

	keys := make([]int, 0, len(mymap))

	for key := range mymap {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return mymap[keys[i]] > mymap[keys[j]]
	})

	result := keys[0:k]

	return result
}
