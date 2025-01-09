package main

import (
	"../issue2"
	"fmt"
)

func main() {
	var param1 = 10
	var param2 = 20

	issue2.SwapPointer(&param1, &param2)

	fmt.Println("After Swap: ", param1, param2)
}
