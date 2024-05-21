package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}

	// onlyGo(arr)
	// withWG(arr)
	// withMutex(arr)
	withBufferedChannel(arr)

}

func onlyGo(arr []int) {

	sum := 0

	for _, num := range arr {
		num := num
		go func() {

			sum += num * num
			fmt.Println(num, sum)
		}()
		time.Sleep(time.Millisecond * 300)

	}

	fmt.Println(sum)
}

func withWG(arr []int) {

	wg := sync.WaitGroup{}
	sum := 0

	for _, num := range arr {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sum += num * num
		}()

	}
	wg.Wait()

	fmt.Println(sum)
}

func withMutex(arr []int) {
	start := time.Now()
	var wg sync.WaitGroup
	var mu sync.Mutex

	var sum int

	for _, num := range arr {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Nanosecond)

			mu.Lock()
			sum += num * num
			mu.Unlock()
		}()
	}
	wg.Wait()

	fmt.Println(sum)

	fmt.Println(time.Now().Sub(start).Seconds())
}

func withChannel(arr []int) {
	start := time.Now()

	ch := make(chan int)
	var sum int

	go func() {
		for _, num := range arr {
			ch <- num * num
		}
		defer close(ch)
	}()

	for {
		v, ok := <-ch
		if !ok {
			break
		}
		sum += v
	}

	fmt.Println(sum)

	fmt.Println(time.Now().Sub(start).Seconds())
}

func withBufferedChannel(arr []int) {
	start := time.Now()

	ch := make(chan int, len(arr))
	var sum int

	go func() {
		for _, num := range arr {
			ch <- num * num
		}
		defer close(ch)
	}()

	for {
		v, ok := <-ch
		if !ok {
			break
		}
		sum += v
	}

	fmt.Println(sum)

	fmt.Println(time.Now().Sub(start).Seconds())
}
