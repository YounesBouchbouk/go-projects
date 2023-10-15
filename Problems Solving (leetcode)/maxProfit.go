package main

import "fmt"

func maxProfit(prices []int) int {

	left, right := 0, 0
	maxProf := 0

	for left < len(prices) {

		if right > len(prices)-1 {
			right = left + 1
			left++

		}

		if left > len(prices)-1 {
			return maxProf
		}

		val := prices[right] - prices[left]

		fmt.Println(fmt.Sprintf("%d - %d = %d", prices[right], prices[left], val))

		if val < 0 {

			left++

		} else {
			if val > maxProf {
				maxProf = val
			}
			right++

		}

	}

	return maxProf

}
