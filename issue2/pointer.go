package issue2

func PointerExample(a int) *int {
	/* 声明实际变量 */
	var ip *int /* 声明指针变量 */

	ip = &a /* 指针变量的存储地址 */

	return ip
	//fmt.Printf("a 变量的地址是: %x\n", &a)

	/* 指针变量的存储地址 */
	//fmt.Printf("ip 变量储存的指针地址: %x\n", ip)

	/* 使用指针访问值 */
	//fmt.Printf("*ip 变量的值: %d\n", *ip)
}

func SwapPointer(a, b *int) {
	*a, *b = *b, *a
}
