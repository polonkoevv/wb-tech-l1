package main

import (
	"fmt"
)

func SetBit(n int64, i int, bitValue int) int64 {
	if bitValue == 1 {
		return n | (1 << i)
	} else {
		return n &^ (1 << i)
	}
}

func main() {
	var n int64 = 42     // Пример числа
	var i int = 3        // Позиция бита для установки (начиная с 0)
	var bitValue int = 0 // Значение бита (1 или 0)

	fmt.Printf("original number: %d (%064b)\n", n, uint64(n))

	bitValue = 0
	n = SetBit(n, i, bitValue)

	fmt.Printf("  result number: %d (%064b)\n\n\n", n, uint64(n))
}
