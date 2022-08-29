package main

import (
	"fmt"
)

var combinations [][]string

func main() {
	example := []string{"1", "2", "3", "4"}
	heapPermutation(example, len(example))
	for _, comb := range combinations {
		fmt.Println(comb)
	}
}

// heap alroithm
func heapPermutation(s []string, size int) {
	// if size becomes 1 then append the result to combinations
	if size == 1 {
		var res = make([]string, 4)
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
