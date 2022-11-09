package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	//slice1()
	//slice2()
	//slice3()
	//slice4()
	//slice5()
	//slice6()
	slice7()
}

func slice7() {
	b := []byte("yufei")
	s := (*string)(unsafe.Pointer(&b))
	fmt.Println(*s)
}

func slice6() {
	s := "雨飞"
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	sh.Cap = sh.Len
	b := (*[]byte)(unsafe.Pointer(sh))
	fmt.Println(string(*b))
}

func slice5() {
	s := "雨飞"
	fmt.Println((*reflect.StringHeader)(unsafe.Pointer(&s)).Data)
	b := []byte(s)
	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&b)).Data)
	s1 := string(b)
	fmt.Println((*reflect.StringHeader)(unsafe.Pointer(&s1)).Data)
}

type slice struct {
	Data uintptr
	Len  int
	Cap  int
}

func slice4() {
	a := [4]string{"yufei", "a", "b", "d"}
	fmt.Printf("%p\n", &a)
	testArray(a)

	b := a[0:]
	fmt.Printf("%d\n", (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data)
	testSlice(b)
}
func testSlice(s []string) {
	fmt.Printf("%d\n", (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data)
}
func testArray(a [4]string) {
	fmt.Printf("%p\n", &a)
}

func slice3() {
	a := [4]string{"yufei", "a", "b", "d"}
	fmt.Println((*slice)(unsafe.Pointer(&a)).Data)
	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&a)).Data)
}

func slice2() {
	a := [4]string{"yufei", "a", "b", "d"}
	s1 := a[0:1]
	s2 := a[:2]

	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&s1)).Data)
	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&s2)).Data)
}

func slice1() {
	a := []string{"yufei", "a"}
	fmt.Println(len(a), cap(a))
	a = append(a, "b", "fd")
	fmt.Println(len(a), cap(a))
}
