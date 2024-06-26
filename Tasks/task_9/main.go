package main

import "fmt"

// записывает числа в канал из массива
func write(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// возводит прочитанные из канала числа и передает в новый канал
func work(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	c := write(2, 3, 5, 10, 45, 7, 5)
	out := work(c)

	// выводим результат из выходного канала
	for v := range out {
		fmt.Println(v)
	}
}
