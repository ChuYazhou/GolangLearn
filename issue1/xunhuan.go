package main

import "fmt"

func main() {

	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Printf("1 到 10 的和是： %d\n", sum)

	// 建立字符串数组
	strings := []string{"hello", "world", "go"}

	for key, value := range strings {
		fmt.Printf("第%d个字符串是%s\n", key+1, value)
	}

	intArray := []int{1, 2, 3, 4, 5}
	for key, value := range intArray {
		fmt.Printf("第%d个数字是%d\n", key+1, value)
	}

	// 建立字典
	maps1 := make(map[int]int)

	for result := 1; result <= 10; result++ {
		maps1[result] = result << 1
		fmt.Printf("maps[%d] = %d\n", result, maps1[result])
	}

	//输出新建的字典
	for key, value := range maps1 {
		fmt.Printf("map1第%d个值是%d\n", key, value)
	}

}
