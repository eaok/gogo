package string

func swapRune(data []rune, x, y int) {
	tmp := data[x]
	data[x] = data[y]
	data[y] = tmp
}

//permutationStr 用递归对字符串中的字符进行全排列
func permutationStr(str []rune, start int) string {
	result := ""
	if str == nil {
		return ""
	}

	if start == len(str)-1 {
		return string(str) + " "
	}

	for i := start; i < len(str); i++ {
		swapRune(str, start, i) //交换start和i所在位置的字符
		result += permutationStr(str, start+1)
		swapRune(str, start, i) //还原start和i所在位置的字符
	}

	return result
}
