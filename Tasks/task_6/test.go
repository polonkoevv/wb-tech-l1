package main

// import (
// 	"context"
// 	"fmt"
// 	"math/rand"
// 	"sync"
// 	"time"
// )

// func writeInChannel(ctx context.Context, wg *sync.WaitGroup, out chan int) {

// 	defer wg.Done()
// 	ticker := time.NewTicker(time.Second)

// 	for {
// 		select {
// 		case <-ctx.Done():
// 			fmt.Println("Stopping sending")
// 			close(out)
// 			return
// 		case <-ticker.C:
// 			v := rand.Int()
// 			out <- v
// 			fmt.Printf("Sent value: %d\n", v)
// 		}

// 	}
// }

// func closeRange(in <-chan int) {
// 	for v := range in {
// 		fmt.Printf("2: Got value: %d\n", v)
// 	}
// 	fmt.Println("\n 2 End of chan")
// }

// func close(in <-chan int) {
// 	for {
// 		v, ok := <-in
// 		if !ok {
// 			fmt.Println("1 Stopping work")
// 		}
// 		fmt.Printf("1: Got value: %d\n", v)
// 	}
// }

// func closeWithContextTimeout(ctx context.Context, wg *sync.WaitGroup, in <-chan int) {
// 	defer wg.Done()
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			fmt.Println("Stopping work")
// 			return
// 		case v := <-in:
// 			fmt.Printf("3: Got value: %d\n", v)
// 		}
// 	}

// }
