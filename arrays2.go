package main

import "fmt"

func main() {
	var x [5]float64
	x[0]=35
	x[1]=58
	x[2]=64
	x[3]=97
	x[4]=31

	var total float64
	total=0
	for i:=0; i<len(x);i++{
		total+=x[i]
	}

	fmt.Println(total/ float64(len(x)))
}
