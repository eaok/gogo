// 高并发限流
// 令牌桶
package main

import (
	"fmt"
	"sync"
	"time"
)

type TokenLimter struct {
	limit float64 // 速率
	burst int     // 桶大小

	mu     sync.Mutex
	tokens float64   // 桶里token数量
	last   time.Time // 上一次小号令牌的时间
}

// 构造函数
func NewTokenLimter(limit float64, burst int) *TokenLimter {
	return &TokenLimter{limit: limit, burst: burst}
}

// 请求是否允许
func (limit *TokenLimter) Allow() bool {
	return limit.AllowN(time.Now(), 1)
}

// 一次请求n个
func (limit *TokenLimter) AllowN(now time.Time, n int) bool {
	limit.mu.Lock()
	defer limit.mu.Unlock()

	// 计算上次补充了多少token
	delta := now.Sub(limit.last).Seconds() * limit.limit
	limit.tokens += delta

	// 如果大于桶则限定到桶的大小
	if limit.tokens > float64(limit.burst) {
		limit.tokens = float64(limit.burst)
	}

	// n 不能大于桶的大小
	if limit.tokens < float64(n) {
		return false
	}

	limit.tokens -= float64(n)
	limit.last = now
	return true
}

// 测试
func main() {
	limiter := NewTokenLimter(3, 5)

	for {
		// 每秒并发4次
		n := 4
		for i := 0; i < n; i++ {
			go func(i int) {
				if !limiter.Allow() {
					fmt.Printf("forbig [%d]\n", i)
				} else {
					fmt.Printf("allow [%d]\n", i)
				}
			}(i)
		}

		time.Sleep(time.Second)
		fmt.Printf("=======================\n")
	}
}
