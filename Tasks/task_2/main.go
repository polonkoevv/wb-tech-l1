package main

import (
	"fmt"
	"time"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}

	OnlyGo(arr)
}

func OnlyGo(arr []int) {
	time.Sleep(time.Second * 5)
	for _, num := range arr {
		num := num
		go func() {
			fmt.Println(num * num)
		}()
	}
}
