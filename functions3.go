package main

import "fmt"

func add (a ...int) int {
	total:=0
	for _, v:= range a {
		total+=v
	}

	return total
}

func main() {
	fmt.Println(add(1,2,3,4,5,6))
	fmt.Println(add(0))
	fmt.Println(add(12,5874,56))
}
