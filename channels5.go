package main

import "fmt"

func main() {
	c:= make(chan int)
	done:= make(chan bool)

	n:=10

	for i:=0; i<n; i++ {
		go func () {
			for j:=0;j<10;j++{
				c<-j
			}
			done<- true
		} ()
	}

	go func () {
		for i:=0;i<n;i++{
			<- done
		}
		close(c)
	} ()

	for t:= range c {
		fmt.Println(t)
	}
}
