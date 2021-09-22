package algorithmThought

//leetcode-cn上可以用到动态规划思想的题
//10. 正则表达式匹配
//62. 不同路径
//70. 爬楼梯

// FibonacciDP 返回斐波那契数列中第n个数
// 动态规划实现
func FibonacciDP(n int) int {
	dp := make([]int, n) // 用于缓存以往结果，以便复用
	dp[0] = 1
	dp[1] = 1

	// 按顺序从小往大算
	for i := 2; i < n; i++ {
		// 使用状态转移方程，同时复用以往结果
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n-1]
}

// FibonacciDP 返回斐波那契数列中第n个数
// 动态规划实现优化版
func FibonacciDPOptimize(n int) int {
	dp := []int{1, 1}

	for i := 2; i < n; i++ {
		dp[0], dp[1] = dp[1], dp[0]+dp[1]
	}

	return dp[1]
}

// FibonacciRE 返回斐波那契数列中第n个数
// 递归实现,为了对比动态规划
func FibonacciRE(n int) int {
	if n < 2 {
		return n
	}

	return FibonacciRE(n-1) + FibonacciRE(n-2)
}
