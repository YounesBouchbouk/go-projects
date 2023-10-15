package main

import "fmt"

func checkIfInclude(m1 map[rune]bool, m2 string, s1 string) bool {
	for _, i := range m2 {
		if m1[i] {
			continue
		}

		return false
	}

	mapR2 := make(map[rune]bool)

	for _, i := range m2 {
		mapR2[i] = true
	}

	fmt.Println(s1)
	fmt.Println(m2)

	fmt.Println(m1)
	fmt.Println(mapR2)

	for _, car := range m2 {
		fmt.Println(car)
		fmt.Println(m1[car])

		if mapR2[car] {
			return false
		}
	}

	return true

}

func checkInclusion(s1 string, s2 string) bool {
	length := len(s1)
	right := 0
	found := false
	mapR := make(map[rune]bool)

	for _, i := range s1 {
		mapR[i] = true
	}

	for i := length; i <= len(s2); i++ {

		s := s2[right:i]

		found = checkIfInclude(mapR, s, s1)
		if found == true {
			return true
		}
		right++
	}

	return false
}
