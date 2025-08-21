package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Чтение данных из канала
func worker(ctx context.Context, num int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d stopped\n", num)
			return
		case c, ok := <-ch:
			if !ok { // Канал закрыт
				fmt.Printf("Worker %d stopped\n", num)
				return
			}
			fmt.Printf("Worker %d got num %d\n", num, c)
		}
	}
}

func main() {
	// Парсинг количества воркеров и времени работы
	n := flag.Int("workers", 1, "number of workers")
	timeout := flag.Int("timeout", 5, "time in seconds before stopping")
	flag.Parse()

	// Создание канала, в который будет реализована постоянная запись данных
	ch := make(chan int)

	// Синхронизация горутин
	var wg sync.WaitGroup

	// Базовый контекст
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Контекст с таймаутом
	ctx, cancel := context.WithTimeout(ctx, time.Duration(*timeout)*time.Second)
	defer cancel()

	// Запуск воркеров
	for i := 1; i <= *n; i++ {
		wg.Add(1)
		go worker(ctx, i, ch, &wg)
	}

	// Постоянная запись данных в канал в отдельной горутине
	go func() {
		for {
			select {
			case <-ctx.Done():
				close(ch)
				return
			default:
				ch <- rand.Intn(100)
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()

	// Ожидание завершения всех горутин
	wg.Wait()
	fmt.Println("All goroutines are stopped")
}
