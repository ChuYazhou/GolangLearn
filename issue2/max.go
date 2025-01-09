package issue2

//func main() {
//	var a, b, c = 100, 200, 300
//
//	fmt.Printf("a,b,c中最大的数字是 %d", max3(a, b, c))
//}

func Max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Max3(a, b, c int) int {
	tem := Max(a, b)
	return Max(tem, c)

}
