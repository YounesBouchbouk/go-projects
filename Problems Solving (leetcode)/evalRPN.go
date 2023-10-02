package main

import (
	"strconv"

	"github.com/YounesBouchbouk/cp/stack"
)

func isOperator(op string) bool {
	if op == "+" || op == "-" || op == "*" || op == "/" {
		return true
	}
	return false
}

func evalRPN(tokens []string) int {

	var mystack stack.Stack

	for _, str := range tokens {
		marks, _ := strconv.Atoi(str)

		if isOperator(str) {
			val1, ok := mystack.Pop()

			if !ok {
				continue
			}
			val2, okk := mystack.Pop()

			if !okk {
				continue
			}

			if str == "+" {
				mystack.Push(val2 + val1)
			} else if str == "*" {
				mystack.Push(val2 * val1)
			} else if str == "/" {
				mystack.Push(val2 / val1)
			} else {
				mystack.Push(val2 - val1)
			}
			continue
		}

		mystack.Push(marks)
	}

	result, err := mystack.Pop()

	if !err {
		return 0
	}
	return result
}
