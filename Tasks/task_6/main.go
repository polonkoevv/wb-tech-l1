package main

import (
	"context"
	"fmt"
	"time"
)

/*
Метод 1: Использование канала для завершения (Done Channel)
Первый метод использует обычное поступление данных в переданный канал. Принцип тот же,
что и при использовании Context. Просто мы сами передаем данные при определенных обстоятельствах
*/
func workerWithDone() {
	fmt.Println("First way")

	ch := make(chan bool)

	go func() {
		for {
			select {
			case <-ch:
				fmt.Println("Goroutine with done channel is shutting down.\n\n")
				return
			default:
				fmt.Println("Working with done channel...")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(3 * time.Second)
	ch <- true
}

/*
Метод 2: Контекст (Context)
Закрытие при помощи context. context дает нам функцию cancel.
Контекст передаем в горутину
Через какое-то время сами вызываем функцию cancel
*/
func workerWithContext() {
	fmt.Println("Second way")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine with context is shutting down.\n\n")
				return
			default:
				fmt.Println("Working with context...")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(3 * time.Second)

}

// Метод 3: Закрытие канала (Closing a Channel)
func workerWithOk() {
	fmt.Println("Third way")

	ch := make(chan int)

	go func() {
		for {
			v, ok := <-ch
			if !ok {
				fmt.Println("Channel closed, goroutine with OK is shutting down.\n\n")
				return
			}
			fmt.Printf("Got value: %d\n", v)
		}
	}()

	for i := 1; i <= 3; i++ {
		ch <- i
		time.Sleep(1 * time.Second)
	}
	close(ch)

}

// Метод 4: Закрытие канала (Closing a Channel) чтение при помощи Range
func workerWithRange() {
	fmt.Println("Fuorth way")

	ch := make(chan int)

	go func() {
		for v := range ch {
			fmt.Printf("Got value: %d\n", v)

		}
		fmt.Println("Channel closed, goroutine with Range is shutting down.\n\n")
		return
	}()

	for i := 1; i <= 3; i++ {
		ch <- i
		time.Sleep(1 * time.Second)
	}
	close(ch)

}

// Метод 5: Таймаут (Timeout)
func workerWithTimeout() {
	fmt.Println("Fifth way")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine with timeout is shutting down due to timeout.\n\n")
				return
			default:
				fmt.Println("Working with timeout...")
				time.Sleep(1 * time.Second)
			}
		}
	}(ctx)

	time.Sleep(4 * time.Second)
	cancel()

}

func main() {

	workerWithDone()
	time.Sleep(1 * time.Second)

	workerWithContext()
	time.Sleep(1 * time.Second)

	workerWithOk()
	time.Sleep(1 * time.Second)

	workerWithRange()
	time.Sleep(1 * time.Second)

	workerWithTimeout()
	time.Sleep(1 * time.Second)

}
