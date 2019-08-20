package numbercrunching

//判断一个数是否是另一个数的平方

//isPower1 直接计算法
func isPower1(n int) bool {
	if n <= 0 {
		println(n, "不是自然数")
		return false
	}

	for i := 1; i < n; i++ {
		m := i * i
		if m == n {
			return true
		} else if m > n {
			return false
		}
	}

	return false
}

//isPower2 二分查找法
func isPower2(n int) bool {
	if n <= 0 {
		println(n, "不是自然数")
		return false
	}

	low, high := 0, n
	for low <= high {
		mid := (low + high) / 2
		power := mid * mid
		if power > n {
			high = mid - 1
		} else if power < n {
			low = mid + 1
		} else {
			return true
		}
	}

	return false
}

//isPower3 减法运算
//利用n^2 = 1 + 3 + ... + 2(n - 1) + 1
func isPower3(n int) bool {
	if n <= 0 {
		println(n, "不是自然数")
		return false
	}

	minus := 1
	for n > 0 {
		n -= minus
		if n == 0 {
			return true
		} else if n < 0 {
			return false
		} else {
			minus += 2
		}
	}

	return false
}
