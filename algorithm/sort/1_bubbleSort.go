package sort

// bobbleSort 冒泡排序
func bubbleSort(array []int) []int {
	if len(array) == 0 {
		return array
	}

	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-1-i; j++ {
			// 升序排序
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}

	return array
}

// bobbleFlagSort 冒泡排序，加如flag改进
func bubbleFlagSort(array []int) []int {
	if len(array) == 0 {
		return array
	}

	flag := 0
	for i := 0; i < len(array)-1; i++ {
		flag = 0
		for j := 0; j < len(array)-1-i; j++ {
			// 升序排序
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				flag = 1
			}
		}
		if flag == 0 {
			break
		}
	}

	return array
}
