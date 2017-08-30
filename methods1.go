package main

import (
	"math"
	"fmt"
)

type circle struct {
	x,y,r float64
}

func (c *circle) area() float64{
	return math.Pi*c.r*c.r
}

func main() {
	c:= circle {0,0, 5}
	fmt.Println(c.area())
}
