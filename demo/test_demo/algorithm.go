package test_demo

// Fibonacci 返回斐波那契数列中第n个数
func Fibonacci(n int) int {
	if n < 2 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}

// Gcd 用辗转相除法求最大公约数
func Gcd(x, y int) int {
	if x < y {
		x, y = y, x
	}

	if y == 0 {
		return x
	}

	return Gcd(y, x % y)
}
