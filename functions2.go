package main

import "fmt"

func f(int) (int,int) {
	return 5,6
}

func main() {
	x, y:= f(3)
	fmt.Println(x,y)
}
