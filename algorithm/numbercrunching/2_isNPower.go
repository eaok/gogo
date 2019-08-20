package numbercrunching

//判断一个数是否为2的n次方

//isNPower1 构造法
//判断1移位后的数值与n是否相等
func isNPower1(n int) bool {
	if n < 0 {
		return false
	}

	for i := 1; i <= n; i <<= 1 {
		if i == n {
			return true
		}
	}

	return false
}

//isNPower2 与操作法
func isNPower2(n int) bool {
	if n < 0 {
		return false
	}

	if n&(n-1) == 0 {
		return true
	}

	return false
}
