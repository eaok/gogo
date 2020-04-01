package test_demo

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestFibonacci 测试Fibonacci函数
func TestFibonacci(t *testing.T) {
	var (
		in       = 7
		expected = 13
	)
	actual := Fibonacci(in)
	if actual != expected {
		t.Errorf("Fibonacci(%d) = %d; expected %d", in, actual, expected)
	}
}

// TestFibonacci2 使用第三方库assert来判断
func TestFibonacci2(t *testing.T) {
	var (
		in       = 7
		expected = 13
	)
	actual := Fibonacci(in)
	assert.Equal(t, actual, expected)
}

// TestGcd 表格驱动来测试Gcd函数
func TestGcd(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "gcd data 1",
			args: args{9, 3},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Gcd(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Gcd() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ExampleFibonacciIterative 用样本测试来测试Fibonacci函数
func ExampleFibonacciIterative() {
	fmt.Println(Fibonacci(4))
	// Output:
	// 3
}

// TestMain 类似于添加了一个中间件
func TestMain(m *testing.M) {
	setUp()
	exitCode := m.Run()
	tearDown()

	os.Exit(exitCode)
}

func setUp() {
	fmt.Println("setUp")
}

func tearDown() {
	fmt.Println("tearDown")
}
