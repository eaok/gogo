package main

import "fmt"

func bigNumberSum(bigNumberA, bigNumberB string) string {
	//1 把两个整数用数组逆序存储，长度为较大整数位加1
	maxLength := 0
	if len(bigNumberA) > len(bigNumberB) {
		maxLength = len(bigNumberA)
	} else {
		maxLength = len(bigNumberB)
	}

	arrayA := make([]int, maxLength + 1)
	arrayB := make([]int, maxLength + 1)
	arrayResult := make([]int, maxLength + 1)
	for i := 0; i < len(bigNumberA); i++ {
		arrayA[i] = int(bigNumberA[len(bigNumberA) - 1 - i] - '0')
	}
	for i := 0; i < len(bigNumberB); i++ {
		arrayB[i] = int(bigNumberB[len(bigNumberB) - 1 - i] - '0')
	}

	//2 遍历数组，按位相加
	for i := 0; i < len(arrayResult); i++ {
		temp := arrayResult[i]
		temp += arrayA[i]
		temp += arrayB[i]
		//判断是否进位
		if temp >= 10 {
			temp -= 10
			arrayResult[i + 1] = 1
		}
		arrayResult[i] = temp
	}

	//3 把arrayResult再次逆序转成string
	findFirst := false
	result := make([]byte, 0)
	for i := len(arrayResult) - 1; i >= 0; i-- {
		if !findFirst {
			if arrayResult[i] == 0 {
				continue
			}
			findFirst = true
		}

		result = append(result, byte(arrayResult[i]) + '0')
	}

	return string(result)
}

func main() {
	fmt.Println(bigNumberSum("426709752318", "95481253129"))
}
