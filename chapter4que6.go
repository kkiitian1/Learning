package main

import "fmt"

func main() {
	var f float64
	fmt.Println("Enter length in feet :")
	fmt.Scanf("%f", &f)
	m:= 0.308*f
	fmt.Println("Length in meters is :", m)
}
