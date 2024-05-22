package main

import (
	"fmt"
	"sync"
)

// Counter - структура счетчика
type MCounter struct {
	value int
	mutex sync.Mutex
}

// Increment - метод для инкрементации счетчика
func (c *MCounter) Increment() {
	c.mutex.Lock()
	c.value++
	c.mutex.Unlock()
}

// Value - метод для получения текущего значения счетчика
func (c *MCounter) Value() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.value
}

func Mutex() {
	var wg sync.WaitGroup
	counter := MCounter{}

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
