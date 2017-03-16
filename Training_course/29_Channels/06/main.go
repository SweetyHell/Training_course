package main

import (
	"fmt"
	"time"
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

	go func() {
		for n := range c {
			fmt.Println(n)
		}
		done <- true
	}()
	go func() {
		for n := range c {
			fmt.Println(n)
		}
		done <- true
	}()
	//you need another go func for this to work
	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(c)
	}()
	time.Sleep(2 * time.Second)
}
