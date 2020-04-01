package sort

//selectSort 选择排序
func selectSort(array []int) {
	for i := 0; i < len(array); i++ {
		tmp := array[i]
		flag := i
		for j := i + 1; j < len(array); j++ {
			if array[j] < tmp {
				tmp = array[j]
				flag = j
			}
		}

		if flag != i {
			array[flag] = array[i]
			array[i] = tmp
		}
	}
}
