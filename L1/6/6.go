package main

////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////

// // Выход по условию
// func main() {
// 	stop := false

// 	go func() {
// 		for !stop {
// 			fmt.Println("Goroutine is running")
// 			time.Sleep(500 * time.Millisecond)
// 		}

// 		fmt.Println("Goroutine is stopped")
// 	}()

// 	time.Sleep(2 * time.Second)
// 	stop = true // Останавливаем горутину
// 	time.Sleep(500 * time.Millisecond)
// }

////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////

// // Через канал уведомления
// func main() {
// 	stop := make(chan struct{})

// 	go func() {
// 		for {
// 			select {
// 			case <-stop:
// 				fmt.Println("Goroutine is stopped")
// 				return
// 			default:
// 				fmt.Println("Goroutine is running")
// 				time.Sleep(500 * time.Millisecond)
// 			}
// 		}
// 	}()

// 	time.Sleep(2 * time.Second)
// 	stop <- struct{}{} // Отправляем сигнал о завершении
// 	time.Sleep(500 * time.Millisecond)
// }

////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////

// // Через контекст
// func main() {
// 	ctx, cancel := context.WithCancel(context.Background())

// 	go func() {
// 		for {
// 			select {
// 			case <-ctx.Done():
// 				fmt.Println("Goroutine is stopped")
// 				return
// 			default:
// 				fmt.Println("Goroutine is running")
// 				time.Sleep(500 * time.Millisecond)
// 			}
// 		}
// 	}()

// 	time.Sleep(2 * time.Second)
// 	cancel() // Останавливаем горутину
// 	time.Sleep(500 * time.Millisecond)
// }

////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////

// // С помощью runtime.Goexit()
// func main() {
// 	start := time.Now()

// 	go func() {
// 		for {
// 			fmt.Println("Goroutine is running")
// 			time.Sleep(500 * time.Millisecond)

// 			if time.Since(start).Seconds() >= 2 { // Останавливаем горутину спустя 2 секунды
// 				fmt.Println("Goroutine is stopped")
// 				runtime.Goexit()
// 			}
// 		}
// 	}()

// 	time.Sleep(time.Second * 3)
// }
