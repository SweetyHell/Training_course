package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	done := make(chan bool)

	go func() {

		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()
	go func() {

		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()
	//you need another go func for this to work
	go func() {
		<-done
		<-done
		close(c)
	}()
	for n := range c {
		fmt.Println(n)
	}

	//time.Sleep(time.Second)
}
