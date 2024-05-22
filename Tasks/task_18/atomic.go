package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Counter - структура счетчика
type ACounter struct {
	value int64
}

// Increment - метод для инкрементации счетчика
func (c *ACounter) Increment() {
	atomic.AddInt64(&c.value, 1)
}

// Value - метод для получения текущего значения счетчика
func (c *ACounter) Value() int64 {
	return atomic.LoadInt64(&c.value)
}

func Atomic() {
	var wg sync.WaitGroup
	counter := ACounter{}

	// Количество горутин
	numGoroutines := 1000

	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	// Ожидаем завершения всех горутин
	wg.Wait()

	fmt.Printf("result counter value: %d\n", counter.Value())
}
