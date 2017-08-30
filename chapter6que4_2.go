package main

import "fmt"

func main() {
	x:= []int {
		48,96,86,68,57,82,63,70,
		37,34,83,27,19,97,9,17}

	min := 1000

	for _, t:= range x {

		if min > t {
			min = t
		}
	}

	fmt.Println("Minimun value in array is:")
	fmt.Println(min)

}
