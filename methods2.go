package main

import (
	"math"
	"fmt"
)

type rectangle struct {
	x1, x2, y1, y2 float64
}

func distance (x1,x2,y1,y2 float64) float64{
	a:= x2-x1
	b:= y2-y1

	return math.Sqrt(a*a + b*b)
}
func (r *rectangle) area () float64 {
	l:= distance(r.x1, r.y1, r.x1, r.y2)
	h:= distance(r.x1, r.y1, r.x2, r.y1)

	return l*h/2
}
func main() {
	r:= rectangle{0,0,10,10}
    fmt.Println(r.area())
}
