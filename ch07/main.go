package main

import (
	"errors"
	"fmt"
)

func main() {
	//testError1()
	//errorTest2()
	//testPanic()
	//testRecover()
	moreDefer()
}
func moreDefer() {
	defer fmt.Println("First defer")
	defer fmt.Println("Second defer")
	defer fmt.Println("Three defer")
	fmt.Println("函数自身代码")
}

func testRecover() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println(p)
		}
	}()
	testPanic()
}

func testPanic() {
	i := -1
	if i < 0 {
		panic("i 小于 0")
	}
}

// 错误嵌套
func errorTest2() {
	me := MyError{errorInfo(), "补充信息"}
	fmt.Println(me.Error())

	// api方式进行错误嵌套
	err := errors.New("原始错误")
	wrapError := fmt.Errorf("额外错误 + %w", err)
	fmt.Println(wrapError)

	// errors.unwrap
	fmt.Println(errors.Unwrap(wrapError))

	// 错误判断：errors.Is方法：相等或者包含target => true
	fmt.Println(errors.Is(wrapError, err))

	var ce *commonError
	// 错误转换:
	fmt.Println(errors.As(err, &ce))
}

type MyError struct {
	err error
	msg string
}

func (myError *MyError) Error() string {
	return myError.err.Error() + myError.msg
}

// 自定义错误&错误断言
func testError1() {
	err := errorInfo()
	if ce, ok := err.(*commonError); ok { // 使用断言将error转换为自定义error
		fmt.Println(ce.Error())
		fmt.Println(ce.errorCode)
	}
}

func errorInfo() error { // 多个返回值之间要用逗号
	return &commonError{
		errorCode: 1,
		msg:       "some error",
	}
}

// 自定义Error
type commonError struct {
	errorCode int
	msg       string
}

// *commonError实现error接口
func (err *commonError) Error() string {
	return err.msg
}
