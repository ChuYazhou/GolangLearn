package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func goTest1() {

	for i := 0; i < 10; i++ {
		fmt.Println("goTest1.....", i)
		time.Sleep(10 * time.Microsecond)
	}

	wg.Done()

}

func goTest2() {

	for i := 0; i < 10; i++ {
		fmt.Println("goTest2.....", i)
		time.Sleep(10 * time.Microsecond)
	}

	wg.Done()

}

func goTest3(num int) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("协程%v正在输出第%v个数据\n", num+1, i+1)
	}
}

func numIsPrime(num int) {
	defer wg.Done()
	var i int
	for i = 2; i < num/2; i++ {
		if num%i == 0 {
			break
		}
	}
	if i > num/i {
		fmt.Println(num, "是素数")
	}
}

func numIsPrimeTest(num int, ch chan int) {
	defer wg.Done()
	var i int
	for i = 2; i < num/2; i++ {
		if num%i == 0 {
			break
		}
	}
	if i > num/i {
		ch <- i
	}
}

func funcCh(ch chan int) {
	defer wg.Done()
	fmt.Printf("%v是素数\n", <-ch)
}
func main() {
	//wg.Add(2)
	//go goTest1()
	//for i := 0; i < 10; i++ {
	//	fmt.Println("main.....", i)
	//	time.Sleep(10 * time.Microsecond)
	//}
	//go goTest2()
	//wg.Wait()
	//cpuNum := runtime.NumCPU()
	//fmt.Println(cpuNum)
	//
	//runtime.GOMAXPROCS(cpuNum - 1)
	//fmt.Println(runtime.NumCPU())
	//for i := 0; i < 10; i++ {
	//	wg.Add(1)
	//	go goTest3(i)
	//}
	//wg.Wait()
	var i int
	// 使用for循环查找素数 并统计程序运行时间
	//time1 := time.Now()
	//for i = 2; i <= 1000000; i++ {
	//	for j = 2; j <= i/j; j++ {
	//		if i%j == 0 {
	//			break
	//		}
	//	}
	//	if j > i/j {
	//		fmt.Println(i)
	//	}
	//}
	//time2 := time.Now()
	//fmt.Printf("程序运行时间：%.2f秒\n", time2.Sub(time1).Seconds())
	// 使用 goroutine 并发查找素数
	var chTest = make(chan int, 1000)
	time1 := time.Now()

	//wg.Add(100)
	for i = 2; i <= 100; i++ {
		//判断i是否为素数
		go numIsPrimeTest(i, chTest)
		go funcCh(chTest)
	}

	time2 := time.Now()

	wg.Wait()

	fmt.Printf("程序运行时间：%.2f秒\n", time2.Sub(time1).Seconds())
	// 使用channel 并发查找素数
}
