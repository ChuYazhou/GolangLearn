package main

import (
	"fmt"
	"sync"
)

func fibonacci(c, quit chan int, wg *sync.WaitGroup) {

	defer wg.Done()

	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y

		case <-quit:
			fmt.Println("quit")

			return
		}
	}
}

func main() {

	wg := &sync.WaitGroup{}
	wg.Add(1)

	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 900
	}()
	go fibonacci(c, quit, wg)

	wg.Wait()
}
