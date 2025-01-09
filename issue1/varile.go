package main

import "fmt"

func main() {
	var x, y int
	var a, b, c = 1, 2, 3
	s := "Hello, world!"
	f := 3.14159265359
	b1 := true

	fmt.Println(x, y, a, b, c, s, f, b1)
	x, y = 10, 20

	fmt.Println(x, y)
	fmt.Printf("x = %d\n, y = %d \n,a, b, c = %d,%d,%d\n\n "+
		"s = %s\n,f = %f \n b1 = %t\n", x, y, a, b, c, s, f, b1)

	x, y = y, x
	fmt.Println(x, y)
}
