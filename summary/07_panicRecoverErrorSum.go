package main

import (
	"errors"
	"fmt"
	"time"
)

// panicExample panic/recover正常使用的例子
func panicExample() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	panic("panic test")
}

// panicOnlyCurrent panic只会触发当前Goroutine的延迟函数调用
func panicOnlyCurrent() {
	defer println("in main")
	go func() {
		defer println("in goroutine")
		panic("panic test")
	}()

	time.Sleep(1 * time.Second)
}

// panicNested panic可以嵌套
func panicNested() {
	defer fmt.Println("in main")
	defer func() {
		defer func() {
			panic("panic again and again")
		}()
		panic("panic again")
	}()

	panic("panic once")
}

// errorNew 使用内置的 errors库生成error信息：
func errorNew() {
	err := errors.New("error")
	if err != nil {
		fmt.Println(err)
	}
}

type ArticleError struct {
	Code    int32
	Message string
}

func (e *ArticleError) Error() string {
	return fmt.Sprintf("[ArticleError] Code=%d, Message=%s", e.Code, e.Message)
}

func NewArticleError(code int32, message string) error {
	return &ArticleError{
		Code:    code,
		Message: message,
	}
}

// errorCustom 生成一个自定义类型的错误
func errorNewCustom() {
	err := NewArticleError(1001, "custom error")
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	//panicExample()
	//panicOnlyCurrent()
	//panicNested()
	//errorNew()
	errorNewCustom()
}
