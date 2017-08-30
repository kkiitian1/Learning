package main

import "fmt"

func main() {
	t:= 1
	for i:=1; i<=100; i++{

		if i%3==0 { t=3}
		if i%5==0 { t=4}
		if i%3==0 && i%5==0 { t=2}

		switch  t{
		case 2: fmt.Println("FizzBuzz")
		case 3: fmt.Println("Fizz")
		case 4: fmt.Println("Buzz")
		default: fmt.Println(i)
		}
		t=1
	}
}
