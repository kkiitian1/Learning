package main

import "fmt"

func main() {
	c:=increment()
	csum:= puller (c)

	for n:=range csum{
		fmt.Println(n)
	}

}
func increment() chan int {
	out:= make(chan int)
	go func () {
		for i := 0; i < 51; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func puller(c chan int) chan int {
	out:= make(chan int)
	go func () {
		var sum int
		for n := range c {
			sum += n
		}

		out <- sum
		close(out)
	}()
	return out
}

