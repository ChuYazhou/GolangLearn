package main

import "fmt"

func dfsdfs(vari uint64) uint64 {
	if vari > 0 {
		return vari * dfsdfs(vari-1)
	}
	return 1
}

func main() {
	start := 15

	fmt.Println(start, "的阶乘是：", dfsdfs(uint64(start)))
}
