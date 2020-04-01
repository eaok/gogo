package sort

//insertSort 插入排序
func insertSort(array []int) {
	if array == nil {
		return
	}

	for i := 1; i < len(array); i++ {
		tmp, j := array[i], i
		if array[j-1] > tmp {
			for j >= 1 && array[j-1] > tmp {
				array[j] = array[j-1]
				j--
			}
		}
		array[j] = tmp
	}
}
