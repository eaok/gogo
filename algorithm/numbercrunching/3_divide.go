package numbercrunching

//不使用除号实现两个正数相除

//divide1 用减法实现
func divide1(m, n int) (res, remain int) {
	res, remain = 0, m

	for m >= n {
		m -= n
		res++
	}
	remain = m

	return
}

//divide2 用移位简化的加法操作
func divide2(m, n int) (res, remain int) {
	res = 0

	for m >= n {
		multi := 1
		for multi*n <= (m >> 1) { //2multi*n > m时结束循环
			multi <<= 1
		}

		res += multi
		m -= multi * n //相减的结果进入下一次循环
	}

	remain = m

	return
}
