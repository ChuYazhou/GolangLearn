package main

import (
	"fmt"
	"sync"
	"time"
)

var wg1 sync.WaitGroup

func writeChan(c chan int) {

	defer wg1.Done()
	for i := 0; i < 10; i++ {
		c <- i
		fmt.Printf("writeChan写入管道%v 数据是%v\n", i, i)
		time.Sleep(2 * time.Second)
	}
	close(c)
}

func readChan(c chan int) {

	defer wg1.Done()
	for value := range c {
		fmt.Printf("从通道%v中读取数据%v\n", value, value)
		time.Sleep(2 * time.Microsecond)
	}
}

func main() {
	ch := make(chan int, 10)
	wg1.Add(1)
	go readChan(ch)
	wg1.Add(1)
	go writeChan(ch)

	wg1.Wait()
	fmt.Println("主线程退出")
}
