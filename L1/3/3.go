package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

// Чтение данных из канала
func worker(num int, ch <-chan int) {
	for c := range ch {
		fmt.Printf("Worker %d got num %d\n", num, c)
	}
}

func main() {
	// Парсинг количества воркеров
	n := flag.Int("workers", 1, "number of workers")
	flag.Parse()

	// Создание канала, в который будет реализована постоянная запись данных
	ch := make(chan int)

	// Запуск воркеров
	for i := 1; i <= *n; i++ {
		go worker(i, ch)
	}

	// Постоянная запись данных в канал
	for {
		ch <- rand.Intn(100)
		time.Sleep(time.Millisecond * 300)
	}
}
