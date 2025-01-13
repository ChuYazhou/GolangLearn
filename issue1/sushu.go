package main

import (
	"../issue2"
	"fmt"
)

func main() {
	arrayInt := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var i, j int
	var a, b, c = 100, 200, 300

	for i = 2; i <= 100; i++ {
		for j = 2; j <= (i / j); j++ {
			if i%j == 0 {
				break
			}
		}
		if j > (i / j) {
			fmt.Printf("%d是素数\n", i)
		}
	}

	fmt.Printf("a,b的最大值是%d\n", issue2.Max(a, b))
	fmt.Printf("a,b,c的最大值是%d\n", issue2.Max3(a, b, c))
	fmt.Printf("数组的平均值是:%d\n", issue2.ArrayAverage(&arrayInt))
	issue2.SetArrayAverage(&arrayInt)
	fmt.Println(arrayInt)
	fmt.Printf("%T\n", arrayInt)
	fmt.Printf("%#v\n", arrayInt)
	fmt.Printf("%+v\n", arrayInt)
	fmt.Printf("%v\n", arrayInt)
	//fmt.Printf("%p\n", arrayInt)
	fmt.Printf("%p\n", &arrayInt)
	fmt.Printf("%p\n", &arrayInt[0])
	fmt.Printf("%p\n", &arrayInt[1])
	fmt.Println(arrayInt)
	tem := issue2.PointerExample(a)
	fmt.Printf("变量地址是：%p\n，变量值是：%d\n", tem, *tem)

	var boolVar int = 20 /* 声明实际变量 */
	var ip, ipp *int     /* 声明指针变量 */

	ip = &boolVar /* 指针变量的存储地址 */

	fmt.Printf("a 变量的地址是: %x\n", &boolVar)

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)
	fmt.Printf("ipp 变量的地址: %x\n", ipp)
	fmt.Printf("ipp 变量的地址: %p\n", ipp)
	fmt.Printf("ipp 变量的地址: %x\n", &ipp)
	fmt.Printf("ipp 变量的地址: %p\n", &ipp)

}
