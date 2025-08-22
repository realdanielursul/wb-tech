package main

import (
	"fmt"
)

// Запись чисел из массива в канал
func gen(nums ...int) <-chan int {
	ch := make(chan int)

	go func() {
		for _, num := range nums {
			ch <- num
		}
		close(ch)
	}()

	return ch
}

// Удваивание чисел в канале
func double(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for i := range in {
			out <- i * i
		}
		close(out)
	}()

	return out
}

func main() {
	ch1 := gen(1, 2, 3, 4, 5)
	ch2 := double(ch1)

	for c := range ch2 {
		fmt.Println(c)
	}
}
