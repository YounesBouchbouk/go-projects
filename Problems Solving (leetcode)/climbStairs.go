package main

func ClimbStairs(n int) int {

	if n < 3 {
		return n
	}

	curr, next := 2, 3

	for i := 3; i < n+1; i++ {
		curr, next = next, curr+next
	}

	return curr
}
