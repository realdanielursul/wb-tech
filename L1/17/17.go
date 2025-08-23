package main

import "fmt"

func BinarySearch(arr []int, value int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := low + (high-low)/2

		if arr[mid] == value {
			return mid
		} else if arr[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func main() {
	fmt.Println(BinarySearch([]int{1, 2, 5, 7, 10, 15}, 5))
}
