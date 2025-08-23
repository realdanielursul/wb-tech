package main

import "fmt"

func quicksort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)/2]
	var left, middle, right []int

	for _, a := range arr {
		if a < pivot {
			left = append(left, a)
		} else if a == pivot {
			middle = append(middle, a)
		} else {
			right = append(right, a)
		}
	}

	res := make([]int, 0, len(arr))
	res = append(res, quicksort(left)...)
	res = append(res, middle...)
	res = append(res, quicksort(right)...)

	return res
}

func main() {
	fmt.Println(quicksort([]int{1, 5, 3, 2, 10, -5, 123, 5}))
}
