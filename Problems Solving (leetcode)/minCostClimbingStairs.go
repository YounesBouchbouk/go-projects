package main

import "fmt"

func minCostClimbingStairs(cost []int) int {
	fmt.Println(cost)

	len := len(cost)
	for i := 2; i < len; i++ {
		cost[i] += minb(cost[i-1], cost[i-2])
	}

	fmt.Println(cost)
	return minb(cost[len-1], cost[len-2])
}

func minb(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
