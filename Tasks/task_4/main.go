package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	var workers int
	var delay time.Duration

	flag.IntVar(&workers, "w", 2, "Number of workers")
	flag.DurationVar(&delay, "t", time.Second, "Delay time")
	flag.Parse()
	fmt.Println(workers)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(workers)

	for i := range workers {
		go worker(ctx, &wg, i+1, ch)
	}

	go func() {
		ticker := time.NewTicker(delay)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				val := rand.Int()
				fmt.Printf("sent value: %d\n", val)
				ch <- val
			case <-ctx.Done():
				return
			}
		}
	}()

	systemChan := make(chan os.Signal, 1)
	signal.Notify(systemChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	sig := <-systemChan
	ctx.Done()
	cancel()
	fmt.Printf("\n\nProgram was stopped with: %s\n", sig.String())
	wg.Wait()

}

func worker(ctx context.Context, wg *sync.WaitGroup, wNum int, ch <-chan int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Context is done. Worker %d stopping\n", wNum)
			return
		case num, ok := <-ch:
			if !ok {
				fmt.Println("Channel is empty. Stopping work")
				return
			}
			fmt.Printf("worker %d: received %d\n", wNum, num)
		}
	}
}
