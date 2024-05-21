package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var delay time.Duration
	var workTime time.Duration
	// Считывание данных для кол-ва ворекров и времени периодичности
	flag.DurationVar(&delay, "t", time.Second, "Delay time")
	flag.DurationVar(&workTime, "wt", time.Second*5, "Working time")
	flag.Parse()

	fmt.Printf("Starting with worktime: %s; delay: %s\n\n", workTime.String(), delay.String())

	ctx, cancel := context.WithTimeout(context.Background(), workTime)
	defer cancel()
	wg := sync.WaitGroup{}
	wg.Add(2)

	ch := make(chan int)

	go worker(&wg, ch)

	go func(ctx context.Context) {
		ticker := time.NewTicker(delay)
		defer ticker.Stop()

		defer wg.Done()
		for {
			select {
			case <-ticker.C:
				v := rand.Int()
				ch <- v
				fmt.Printf("Sent value: %d\n", v)
			case <-ctx.Done():
				close(ch)
				return
			}
		}

	}(ctx)

	wg.Wait()

}

func worker(wg *sync.WaitGroup, in <-chan int) {
	for v := range in {
		fmt.Printf("Received value: %d\n", v)
	}
	defer wg.Done()
	fmt.Println("\n\nWorker stopped")
}
