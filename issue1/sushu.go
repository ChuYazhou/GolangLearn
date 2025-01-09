package main

import (
	"../issue2"
	"fmt"
)

func main() {
	var i, j int
	var a, b = 100, 200

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
}
