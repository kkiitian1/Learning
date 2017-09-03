package main

import (
//	"sync"
	"fmt"
//	"sync"
)

func main() {
	c:= make (chan int)
	done:= make(chan bool)

	go func() {
		for i:=0; i<10; i++ {
			c<- i
		}
		//wg.Done()
		done<- true
	}()

	go func() {
		for i:=0; i<10; i++ {
			c<- i
		}
		done<- true
		//wg.Done()
	}()

	go func () {
		//wg.Wait()
		<-done
		<-done
		close(c)
	} ()

	for n:= range c {
		fmt.Println(n)
	}

}

