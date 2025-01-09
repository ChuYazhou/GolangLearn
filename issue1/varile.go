package main

import "fmt"

const (
	test  = "abc"
	test2 = len(test)
	test3 = iota
	test4 = iota
	test5 = "haha"
	test6
	test7 = iota
	test8 = 100
)

func main() {
	a, b := 123, "hello"
	var x, y int
	var c = 1
	s := "Hello, world!"
	f := 3.14159265359
	var b1 bool = true

	fmt.Println(x, y, a, b, c, s, f, b1)
	x, y = 10, 20

	fmt.Printf("x = %d,y = %d\n", x, y)

	//fmt.Printf("x = %d\n, y = %d \n,a, b, c = %d,%d,%d\n\n "+
	//	"s = %s\n,f = %f \n b1 = %t\n", x, y, a, b, c, s, f, b1)
	fmt.Println(&x, &y)
	swapping(&x, &y)
	fmt.Printf("x = %d,y = %d\n", x, y)
	fmt.Println(&x, &y)
	_, sub := number(x, y)
	fmt.Printf("x+y = %d,x-y = %d\n", sub, sub)
}

func swapping(a, b *int) {
	*a, *b = *b, *a
}

func number(a, b int) (int, int) {
	return a + b, a - b
}
