package main

import "fmt"

func SetBitTo1(num int64, index int) int64 {
	num = num | (1 << index)
	return num
}

func SetBitTo0(num int64, index int) int64 {
	num = num &^ (1 << index)
	return num
}

func main() {
	fmt.Println(SetBitTo0(5, 0)) // 5 -> 4
	fmt.Println(SetBitTo1(4, 0)) // 4 -> 5
}
