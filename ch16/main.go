package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	//unsafe1()
	//unsafe2()
	unsafe3()
}

type person struct {
	name string
	age  int
}

func unsafe3() {
	bo := true
	fmt.Println(unsafe.Sizeof(bo))
	fmt.Println(&bo)
	i := int8(32)
	fmt.Println(unsafe.Sizeof(i))
	fmt.Println(&i)
	i16 := int16(345)
	fmt.Println(unsafe.Sizeof(i16))
	fmt.Println(&i16)
	fmt.Println(unsafe.Sizeof(int32(23)))
	fmt.Println(unsafe.Sizeof(int64(23)))
	fmt.Println(unsafe.Sizeof(int(23)))
	fmt.Println(unsafe.Sizeof(string("雨飞")))
	fmt.Println(unsafe.Sizeof([]string{"yufei", "a"}))
	fmt.Println(reflect.TypeOf([]string{"yufei", "yufei2"}).Kind())
}

func unsafe2() {
	p := new(person)
	np := (*string)(unsafe.Pointer(p))
	*np = "雨飞"

	ap := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + unsafe.Offsetof(p.age)))
	*ap = 28
	fmt.Println(*p)
}

func unsafe1() {
	i := 10
	ip := &i
	var fp *float64 = (*float64)(unsafe.Pointer(ip))
	*fp = *fp * 3
	fmt.Println(i)
}
