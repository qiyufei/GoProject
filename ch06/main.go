package main

import "fmt"

func main() {
	//testStruct()
	//testInterface()
	test3()
}

func test3() {
	p := person{
		name: "雨飞",
		age:  18,
		addr: address{
			city:     "佳木斯",
			province: "黑龙江",
		},
	}
	// 值类型接收者实现接口
	printStringerEp(p)
	// 指针类型接收者实现接口
	printStringerEp(p)
}

func testInterface() {
	addr := address{
		city:     "佳木斯",
		province: "黑龙江",
	}
	p := person{
		name: "雨飞",
		age:  18,
		addr: addr,
	}
	printStringerEp(p)
	printStringerEp(addr)
}

// 面向接口编程
func printStringerEp(ep StringerEp) {
	ep.String()
}

// 让address类实现接口 <=> 让address类接收方法
func (addr address) String() int {
	fmt.Printf("city is %s, province is %s\n", addr.city, addr.province)
	return 0
}

// 让person类实现接口 <=> 让person类接收方法
func (p person) String() int {
	fmt.Printf("name is %s, age is %d\n", p.name, p.age)
	return 0
}

type StringerEp interface {
	String() int
}

func testStruct() {
	// 创建并初始化结构体
	p := person{ //字面量初始化用大括号
		name: "雨飞",
		age:  18,
		addr: address{
			city:     "佳木斯",
			province: "黑龙江",
		}, //最后这里跟着一个逗号
	}
	fmt.Println(p.name, p.addr.city)
}

// 定义结构体
type person struct {
	name string
	age  int
	addr address
}

type address struct {
	city     string
	province string
}
