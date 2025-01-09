package main

import "fmt"

func main() {

	//创建切片
	var slice1 = make([]int, 5, 6)
	slice1 = []int{1, 2, 3, 5, 4}

	// 切片增加
	slice1 = append(slice1, 6) // 切片增加

	// 切片扩容

	slice2 := make([]int, len(slice1), 2*cap(slice1))
	slice2 = append(slice2, slice1[0], slice1[1], slice1[2], slice1[3], slice1[4], slice1[5])
	// 切片复制
	copy(slice2, slice1)

	fmt.Println(slice1)
	fmt.Println(slice2)

	var string1 = "hello world"
	var string2 = "go"

	for k, v := range string1 {
		fmt.Println(k, v)
		fmt.Printf("index:%d , value : %c\n", k, v)
	}

	for k, v := range string2 {
		fmt.Println(k, v)
		fmt.Printf("index:%d , value : %c\n", k, v)
	}

}
