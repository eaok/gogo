package main

import (
	"fmt"
)

//removeDigits 删除整数的k个数字，获得删除后的最小值
func removeKDigits(num string, k int) string {
	numNew := num
	for i := 0; i < k; i++ {
		hasCut := false
		//从左向右遍历，找到比自己右侧大的数字并删除
		for j := 0; j < len(numNew) - 1; j++ {
			if numNew[j] > numNew[j + 1] {
				numNew = numNew[:j] + numNew[j + 1:]
				hasCut = true
				break
			}
		}

		//如果没有找到，就删除最后一个数字
		if !hasCut {
			numNew = numNew[:len(numNew) - 1]
		}

		//清除整数左侧的数字0
		if numNew[0] == '0' {
			numNew = numNew[1:]
		}

		//如果整数的所有数字都被删除了
		if len(numNew) == 0 {
			return "0"
		}
	}

	return numNew
}

//removeKDigitsV2 使用栈来优化
func removeKDigitsV2(num string, k int) string {
	newLength := len(num) - k

	//创建一个栈，用于接收所有的数字
	stack := make([]byte, len(num))
	top := 0
	for i := 0; i < len(num); i++ {
		//当栈顶数字大于遍历到的数字时出栈，相当于删除数字
		for top > 0 && stack[top - 1] > num[i] && k > 0 {
			top -= 1
			k -= 1
		}
		stack[top] = num[i]
		top++
	}

	//找到栈中第一个非零的数字，构建新的字符串
	offset := 0
	for offset < newLength && stack[offset] == '0' {
		offset++
	}

	if offset == newLength {
		return "0"
	}

	return string(stack[offset:])
}

func main() {
	fmt.Println(removeKDigits("1593212", 3))
	fmt.Println(removeKDigitsV2("1593212", 3))
	fmt.Println(removeKDigits("30200", 1))
	fmt.Println(removeKDigitsV2("30200", 1))
	fmt.Println(removeKDigits("10", 2))
	fmt.Println(removeKDigitsV2("10", 2))
	fmt.Println(removeKDigits("541270936", 2))
	fmt.Println(removeKDigitsV2("541270936", 2))
}
