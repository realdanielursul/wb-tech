package main

import "fmt"

func main() {
	x := 7
	y := 15

	// Обмен значений
	x = x + y
	y = x - y
	x = x - y

	fmt.Println(x, y)
}
