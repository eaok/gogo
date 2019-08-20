package main

import (
	"fmt"
	"log"
)

func findNearestNumber(numbers []int) []int {
	//1 从后向前找到逆序区域的前一位
	index := 0
	for i := len(numbers) - 1; i > 0; i-- {
		if numbers[i] > numbers[i - 1] {
			index = i
		}
	}
	if index == 0 {
		return nil
	}

	//2 把逆序区域的前一位和逆序区域中刚刚大于它的数字交换位置
	resultNumbers := make([]int, len(numbers))
	copy(resultNumbers, numbers)
	head := resultNumbers[index - 1]
	for i := len(resultNumbers) - 1; i > 0; i-- {
		if head < resultNumbers[i] {
			resultNumbers[index - 1] = resultNumbers[i]
			resultNumbers[i] = head
			break
		}
	}

	//3 把原来的逆序区域转为顺序
	for i, j := index, len(resultNumbers) - 1; i < j; i++ {
		temp := resultNumbers[i]
		resultNumbers[i] = resultNumbers[j]
		resultNumbers[j] = temp
		j--
	}
	log.Println(index, numbers, resultNumbers)

	return resultNumbers
}

func main() {
	numbers := []int{5, 4, 2, 3, 1}

	for i := 0; i < 2; i++ {
		numbers = findNearestNumber(numbers)
		fmt.Println(numbers)
	}
}
