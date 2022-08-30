package main

import (
	"fmt"
)

var combinations [][]int

func main() {
	example := []int{1, 2, 3, 4}
	checkTotalValidTimes(example)
}

func checkTotalValidTimes(s []int) {
	if len(s) != 4 {
		panic("4 integers are required")
	}

	for _, v := range s {
		if v < 0 || v > 9 {
			panic("Integers has to be between 0 and 9")
		}
	}

	heapPermutation(s, len(s))
	totalValidTimes := 0
	for _, comb := range combinations {
		isValid := isValidTime(comb)
		fmt.Println(comb, isValid)

		if isValid {
			totalValidTimes++
		}
	}

	fmt.Println("Total valid times: ", totalValidTimes)
}

// heap algorithm
func heapPermutation(s []int, size int) {
	// if size becomes 1 then append the result to combinations
	if size == 1 {
		var res = make([]int, 4)
		copy(res, s)
		combinations = append(combinations, res)
		return
	}

	for i := 0; i < size; i++ {
		heapPermutation(s, size-1)
		// if size is odd, swap first and last element
		// if size is even, swap ith and last element
		if size%2 != 0 {
			s[0], s[size-1] = s[size-1], s[0]
		} else {
			s[i], s[size-1] = s[size-1], s[i]
		}

	}
}

func isValidTime(s []int) bool {
	// check hours
	if s[0] > 2 {
		return false
	}

	if s[0] == 2 && s[1] > 3 {
		return false
	}

	// check minutes
	if s[2] > 5 {
		return false
	}

	if s[3] > 9 {
		return false
	}

	return true
}
