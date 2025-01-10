package main

import (
	"fmt"
	"sync"
)

// 创建一个函数 读取切片数据，并载入到通道中
func sum(slice1 []int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	sum := 0
	for _, value := range slice1 {
		sum += value
	}
	ch <- sum
	fmt.Printf("The sum is: %d\n", sum)

}

func main() {
	// 创建一个切片
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// 创建一个通道
	ch := make(chan int, 2)
	var wg sync.WaitGroup

	wg.Add(2)

	go sum(slice1[:len(slice1)/2], ch, &wg)
	go sum(slice1[len(slice1)/2:], ch, &wg)

	wg.Wait()

	x := <-ch
	fmt.Println(x)
	y := <-ch
	fmt.Println(y)
	close(ch)
}
