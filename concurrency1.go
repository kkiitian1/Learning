package main

import (
	"fmt"
	"sync"
	"time"
	"runtime"
)

//run -race main.go
var wg sync.WaitGroup

//func init () {
//	runtime.GOMAXPROCS(runtime.NumCPU())
//}

func foo () {
	for i:=0; i<50; i++ {
		fmt.Println("FOO :", i)
		time.Sleep(time.Duration(5*time.Millisecond))
	}
	wg.Done()
}

func bar () {
	for i:=100; i<150;i++ {
		fmt.Println("BAR :", i)
		time.Sleep(time.Duration(5*time.Millisecond))
	}
	wg.Done()
}
func main() {

	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
	//go run -race main.go

}

