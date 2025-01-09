package issue2

func ArrayAverage(param *[]int) int {

	// 获取数组长度
	length := len(*param)
	var tem int

	// 获取数组总和
	for _, value := range *param {
		tem += value
	}
	// 计算平均值
	return int(tem / length)

}

func SetArrayAverage(param *[]int) {
	average := ArrayAverage(param)
	*param = append(*param, average)
}
