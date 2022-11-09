package main

import (
	"errors"
	"fmt"
)

func main() {
	//fmt.Println(sum(1, 2))

	//result, err := sum1(1, 1)
	//if err != nil {
	//	fmt.Println(result)
	//} else {
	//	fmt.Println(err)
	//}

	//result, _ := sum2(-1, 3)
	//fmt.Println(result)

	//fmt.Println(sum3(1, 9, 3))

	//anonymousFunc()

	//testClosure()

	age := Age(25)
	age.string()     //打印 30
	fmt.Println(age) // 打印25

	age.string2()
	fmt.Println(age)
}

type Age uint

// 指针型接受者
func (age *Age) string2() {
	*age = 30
}

// 值型接收者
func (age Age) string() {
	age = 30 // 不会修改调用者的值
	fmt.Println(age)
}

func testClosure() {
	// 调用closure()返回一个函数变量
	closureFunc := closure()
	// 调用这个函数，获取函数返回值，该值在closure内部定义的，被闭包持有，被闭包累增
	i := closureFunc()
	fmt.Println(i)    // 打印
	i = closureFunc() // 再调用一次，该函数变量持有的i再次加一
	fmt.Println(i)
	i = closureFunc() // 再调用一次，该函数变量持有的i再次加一
	fmt.Println(i)

	// 简化上面的
	// 一次调用closure()获取一个函数变量 closureFunc
	// 同一个 closureFunc 函数变量，一直持有自己的i，那么调用这个函数变量，其持有的i就会累加
	closureFunc2 := closure()
	fmt.Println(closureFunc2())
	fmt.Println(closureFunc2())
	fmt.Println(closureFunc2())
}

// 定义一个闭包/嵌套函数/函数中定义函数
func closure() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// 匿名函数
func anonymousFunc() {
	// 定义一个匿名函数并且调用
	sumFunc := func(a, b int) int {
		return a + b
	}
	fmt.Println(sumFunc(1, 2))
}

// 函数参数是可变的，不确定数量的
func sum3(params ...int) int { // 可变参数的底层是切片 slice
	sum := 0
	for _, i := range params { //for range return index and value
		sum += i
	}
	return sum
}

// 返回参数可以被命名
func sum2(a, b int) (result int, err error) {
	if a < 0 || b < 0 {
		result = 0
		err = errors.New("less than 0")
	} else {
		result = a + b
		err = nil
	}
	return // 对返回参数赋值后，这里直接return即可
}

func sum1(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("a, b can not less than zero")
	} else {
		return a + b, nil
	}
}

func sum(a, b int) int {
	return a + b
}
