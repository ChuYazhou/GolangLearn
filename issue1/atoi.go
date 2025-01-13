package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string = "123"

	if test1, err := strconv.Atoi(str); err == nil {
		fmt.Println(test1)
	}
}
