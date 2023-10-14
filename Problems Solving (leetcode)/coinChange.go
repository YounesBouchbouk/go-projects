package main

import (
	"fmt"
	"sort"
)

// this approche can't work thank you
func coinChange(coins []int, amount int) int {

	lenOfcoins := len(coins) - 1
	var results []int

	sort.Ints(coins)

	if amount == 0 {
		return 0
	}

	for lenOfcoins >= 0 {
		if amount-coins[lenOfcoins] > 0 {
			amount = amount - coins[lenOfcoins]
			results = append(results, amount)
		} else if amount-coins[lenOfcoins] < 0 {
			lenOfcoins--
			continue
		} else {
			results = append(results, 0)
			break
		}

	}

	fmt.Println(results)

	if len(results) > 0 && results[len(results)-1] == 0 {
		return len(results)
	}

	return -1
}
