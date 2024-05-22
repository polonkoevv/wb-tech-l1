package main

import (
	"fmt"
	"sync"
)

// Counter - структура счетчика
type Counter struct {
	value int
	mutex sync.Mutex
}

// Increment - метод для инкрементации счетчика
func (c *Counter) Increment() {
	c.mutex.Lock()
	c.value++
	c.mutex.Unlock()
}

// Value - метод для получения текущего значения счетчика
func (c *Counter) Value() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.value
}

func Classic() {
	var wg sync.WaitGroup
	counter := Counter{}

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

func main() {

	// Classic()
	// Atomic()
	Mutex()

}
