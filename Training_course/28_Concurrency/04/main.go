package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var mutex sync.Mutex
var counter int64

func main() {
	wg.Add(2)
	go incrementor("Foo:")
	go incrementor("Bar:")
	wg.Wait()
	fmt.Println("Final counter:", counter)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
		//mutex.Lock()
		//counter++
		atomic.AddInt64(&counter, 1)
		fmt.Println(s, i, "Counter:", counter)
		//mutex.Unlock()
	}
	wg.Done()
}
