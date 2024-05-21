package main

import (
	"fmt"
	"polonkoevv/wb-l1/Tasks/task_24/point"
)

func main() {
	var x1 float64
	var y1 float64

	var x2 float64
	var y2 float64

	fmt.Print("Set value for point x1: ")
	fmt.Scanf("%f\n", &x1)
	fmt.Print("Set value for point y1: ")
	fmt.Scanf("%f\n", &y1)
	fmt.Print("Set value for point x2: ")
	fmt.Scanf("%f\n", &x2)
	fmt.Print("Set value for point y2: ")
	fmt.Scanf("%f\n", &y2)

	p1 := point.New(x1, y1)
	p2 := point.New(x2, y2)

	fmt.Printf("distance between (%f, %f) and (%f, %f) is %f\n", x1, y1, x2, y2, point.Distance(p1, p2))
}
