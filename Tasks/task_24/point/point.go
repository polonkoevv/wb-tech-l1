package point

import "math"

type Point struct {
	x float64
	y float64
}

func New(x float64, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

func Distance(p1 Point, p2 Point) float64 {
	return math.Sqrt((p2.x-p1.x)*(p2.x-p1.x) + (p2.y-p1.y)*(p2.y-p1.y))
}
