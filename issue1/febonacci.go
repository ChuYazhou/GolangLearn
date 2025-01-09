package main

import "fmt"

func fabona(n int) int {
	if n < 2 {
		return n
	}
	return fabona(n-1) + fabona(n-2)
}

func main() {

	for i := 0; i < 20; i++ {
		fmt.Printf("%d ", fabona(i))
	}
}
