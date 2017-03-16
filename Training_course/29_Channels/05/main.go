package main

import (
	"fmt"
)

func main() {
	n := 10
	c := make(chan int)
	done := make(chan bool)

	for i := 0; i < n; i++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
			done <- true
		}()
	}

	//you need another go func for this to work
	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}

	//time.Sleep(time.Second)
}
