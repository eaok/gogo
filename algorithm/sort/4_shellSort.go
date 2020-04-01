package sort

//shellSort 希尔排序
func shellSort(array []int) []int {
	//外层步长控制
	for step := len(array) / 2; step > 0; step /= 2 {
		//开始插入排序
		for i := step; i < len(array); i++ {
			//满足条件则插入
			for j := i - step; j >= 0 && array[j+step] < array[j]; j -= step {
				array[j], array[j+step] = array[j+step], array[j]
			}
		}
	}

	return array
}
