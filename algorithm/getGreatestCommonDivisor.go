package main

import (
	"fmt"
)

//gcd1 使用枚举
func gcd1(a, b int) int {
	if a > b {
		a, b = b, a
	}
	small, big := a, b
	if big % small == 0 {
		return small
	}

	for i := small / 2; i > 1; i-- {
		if small % i == 0 && big % i == 0 {
			return i
		}
	}

	return 1
}

//gcd2 使用辗转相除法
func gcd2(a, b int) int {
	if a > b {
		a, b = b, a
	}
	small, big := a, b

	if big % small == 0 {
		return small
	}

	return gcd2(big % small, small)
}

//gcd3 使用更相减损术
func gcd3(a, b int) int {
	if a == b {
		return a
	}
	if a > b {
		a, b = b, a
	}
	small, big := a, b

	return gcd3(big - small, small)
}

//gcd4 使用更相减损术和移位相结合
func gcd4(a, b int) int {
	if a == b {
		return a
	}
	if (a & 1 == 0) && (b & 1 == 0) {
		return gcd4(a >> 1, b >> 1) << 1
	} else if (a & 1 == 0) && (b & 1 != 0) {
		return gcd4(a >> 1, b)
	} else if (a & 1 != 0) && (b & 1 == 0) {
		return gcd4(a, b >> 1)
	} else {
		if a > b {
			a, b = b, a
		}
		small, big := a, b

		return gcd4(big - small, small)
	}
}

func main() {
	fmt.Println(gcd1(55, 110))
	fmt.Println(gcd2(55, 110))
	fmt.Println(gcd3(55, 110))
	fmt.Println(gcd4(55, 110))
}
