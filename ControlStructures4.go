package main

import "fmt"

func main() {
	i := 3
	//fmt.Println("Enter a number :")
	//fmt.Scanf("%f", &i)
	switch i {
	case 0: fmt.Println("zero")
	case 1: fmt.Println( "one")
	case 2: fmt.Println("two")
	case 3: fmt.Println("three")
	case 4: fmt.Println("four")
	case 5: fmt.Println(" five")
	default: fmt.Println("Unknown Number")
	}
}