package main

import "fmt"

func main() {
	xs:= []float64 {98,93,77,82,83}
	fmt.Println(average(xs))
}

func average (x []float64) float64 {
	total:= 0.0
	for _, t:= range x {
		total+=t
	}

	return total/float64(len(x))
}
