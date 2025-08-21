package main

import (
	"fmt"
	"sync"
	"time"
)

// Реализация конкурентной мапы
type ConcurrentMap[T1 comparable, T2 any] struct {
	mp map[T1]T2
	mu sync.Mutex
}

func NewConcurrentMap[T1 comparable, T2 any]() *ConcurrentMap[T1, T2] {
	return &ConcurrentMap[T1, T2]{
		mp: make(map[T1]T2),
		mu: sync.Mutex{},
	}
}

func (cm *ConcurrentMap[T1, T2]) Set(key T1, value T2) {
	cm.mu.Lock()
	cm.mp[key] = value
	cm.mu.Unlock()
}

func (cm *ConcurrentMap[T1, T2]) Get(key T1) (T2, bool) {
	cm.mu.Lock()
	value, ok := cm.mp[key]
	cm.mu.Unlock()
	return value, ok
}

func main() {
	start := time.Now()
	var wg sync.WaitGroup

	mp := NewConcurrentMap[int, bool]()

	// Запись 100 единиц данных (каждая запись длится 1 сек)
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mp.Set(i, true)
			time.Sleep(time.Second)
		}()
	}

	wg.Wait()

	fmt.Println(mp.mp)             // Содержимое мапы от 1 до 100
	fmt.Println(time.Since(start)) // 1.00130224s вместо 100s, следовательно запись производится конкурентно
}
