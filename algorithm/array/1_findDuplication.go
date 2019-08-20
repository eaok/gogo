package array

//findDupByHash hash法
//再数组中找到唯一重复的元素
func findDupByHash(arr []int) int {
	if arr == nil {
		return -1
	}

	data := map[int]bool{}
	for _, v := range arr {
		if _, ok := data[v]; ok {
			return v
		}
		data[v] = true
	}

	return -1
}

//findDupByXOR 异或法
func findDupByXOR(arr []int) int {
	if arr == nil {
		return -1
	}

	result := 0
	for _, v := range arr {
		result ^= v
	}
	for i := 1; i < len(arr); i++ {
		result ^= i
	}

	return result
}

//findDupByMap 数据映射法
func findDupByMap(arr []int) int {
	if arr == nil {
		return -1
	}

	index := 0

	for {
		if arr[index] >= len(arr) { //数组中元素的值只能小于len,否则会溢出
			return -1
		}
		if arr[index] < 0 { //找到重复的数
			break
		}
		arr[index] *= -1 //访问过后就取反标记
		index = arr[index] * -1
		if index >= len(arr) {
			return -1 //数组中有非法数字
		}
	}

	return index
}

//findDupByLoop 环形相遇法
//相当于找入环点
func findDupByLoop(arr []int) int {
	if arr == nil {
		return -1
	}

	slow, fast := 0, 0
	for ok := true; ok; ok = slow != fast {
		fast = arr[arr[fast]]
		slow = arr[slow]
	}
	fast = 0
	for ok := true; ok; ok = slow != fast {
		fast = arr[fast]
		slow = arr[slow]
	}

	return slow
}
