package main

import (
	"sync"
	"fmt"
	"time"
	"math/rand"
	"sync/atomic"
)

var wg sync.WaitGroup
//var mutex sync.Mutex
var counter int64

func incrementor (s string) {

	for i:=0; i<20; i++ {
		time.Sleep(time.Duration(rand.Intn(20))*time.Millisecond)
		//mutex.Lock()
		//x:= counter
		//x++
		//counter++
		atomic.AddInt64(&counter, 1)
		fmt.Println(s, i, "counter :", counter )
		//mutex.Unlock()
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go incrementor ("FOO :")
	go incrementor ("BAR :")
	wg.Wait()
	fmt.Println("Final counter is :", counter)
}