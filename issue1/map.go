package main

import "fmt"

func main() {

	// 初始化字典

	var dict map[string]string

	dict = make(map[string]string, 10)

	// 读取字典

	dict["中国"] = "北京"
	dict["日本"] = "东京"

	for key := range dict {
		fmt.Println(key, "的首都是", dict[key])
	}

	cap1, ok := dict["美国"]
	if ok {
		fmt.Printf("美国的首都是：%s\n", cap1)
	} else {
		fmt.Println("美国不在该站点")
	}

	delete(dict, "日本")

}
