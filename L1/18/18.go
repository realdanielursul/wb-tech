package main

import (
	"fmt"
	"sync"
)

type Incrementer struct {
	counter int
	mx      sync.Mutex
}

func (i *Incrementer) Increment(wg *sync.WaitGroup) {
	defer wg.Done()
	i.mx.Lock()
	i.counter++
	i.mx.Unlock()
}

func main() {
	var inc Incrementer
	var wg sync.WaitGroup

	wg.Add(100000)
	for i := 0; i < 100000; i++ {
		go inc.Increment(&wg)
	}

	wg.Wait()
	fmt.Println(inc.counter) // 100000
}
