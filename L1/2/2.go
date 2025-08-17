package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	nums := []int{2, 4, 6, 8, 10}
	for _, n := range nums {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			fmt.Printf("%d^2 = %d\n", num, num*num)
		}(n)
	}

	wg.Wait()
}
