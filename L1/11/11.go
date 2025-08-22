package main

import (
	"fmt"
	"slices"
)

func FindIntersection(arr1, arr2 []int) []int {
	ans := make([]int, 0)

	for _, value := range arr1 {
		if slices.Contains(arr2, value) {
			ans = append(ans, value)
		}
	}

	return ans
}

func main() {
	fmt.Println(FindIntersection([]int{1, 2, 3}, []int{2, 3, 4})) // [2, 3]
}
