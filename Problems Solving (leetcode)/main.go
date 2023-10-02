package main

import (
	"fmt"

	linkedlist "github.com/YounesBouchbouk/cp/linkedList"
)

// func main() {

// 	// fmt.Println("topKFrequent: ", topKFrequent([]int{2}, 1))
// 	// fmt.Println("productExceptSelf: ", ProductExceptSelf([]int{1, 2, 3, 4}))
// 	// fmt.Println("longest : ", longestConsecutive([]int{0, 0, -1}))
// 	// fmt.Println("lengthOfLongestSus[left]string : ", lengthOfLongestSus[left]string("pwwkew"))
// 	// fmt.Println("is isPalindrome : ", isPalindrome("A man, a plan, a canal: Panama"))
// 	// fmt.Println("threeSum : ", threeSum([]int{-1, 0, 1, 2, -1, -4}))
// 	// fmt.Println("nb trap : ", trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
// 	// fmt.Println("isValid : ", isValid("({})"))
// 	// fmt.Println("evalRPN : ", evalRPN([]string{"4", "13", "5", "/", "+"}))
// 	// fmt.Println("searchMatrix : ", searchMatrix([][]int{
// 	// 	{1, 3, 5, 7},
// 	// 	{10, 11, 16, 20},
// 	// 	{23, 30, 34, 60},
// 	// }, 3))
// 	// fmt.Println("the min value : ", searchRotatedSorted([]int{1, 3}, 0))

// }

func main() {
	list := &linkedlist.List{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)

	fmt.Println("Initial List: ")
	list.PrintList()

	list.Reverse()
	fmt.Println("List reverse  ")
	list.PrintList()
}
