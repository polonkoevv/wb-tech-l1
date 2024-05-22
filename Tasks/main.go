package main

import "fmt"

type Number interface {
	int | float64
}

func Sum[V Number](a, b V) V {
	return a + b
}

func main() {
	a, b := 1, 2
	fmt.Println(Sum[int](a, b)) // используем Sum для Int

	c, d := 1.2, 3.4
	fmt.Println(Sum[float64](c, d)) // используем Sum для Floats
}
