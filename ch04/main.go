package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//arrayTest()
	//sliceTest()
	//mapTest()
	//byteStrTest()
	array2()
}

func array2() {
	aa := [3][3]int{}

	aa[0][0] = 1

	aa[0][1] = 2

	aa[0][2] = 3

	aa[1][0] = 4

	aa[1][1] = 5

	aa[1][2] = 6

	aa[2][0] = 7

	aa[2][1] = 8

	aa[2][2] = 9

	fmt.Println(aa)
}

func byteStrTest() {
	s := "hello雨飞"
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Println(s[0], s[2], s[8])
	fmt.Println(utf8.RuneCountInString(s))
	for i, c := range s {
		fmt.Println(i, c)
	}
}

func mapTest() {
	// 字面量创建
	map2 := map[string]int{"忍": 1}
	i := map2["忍"]
	fmt.Println(i)
	delete(map2, "忍")
	fmt.Println(map2)

	// 创建map
	map1 := make(map[string]int)
	map1["雨飞"] = 18
	fmt.Println(map1)

	//判断key是否存在
	j, ok := map1["加油"]
	if ok {
		fmt.Println(j)
	}

	map1["你好"] = 2
	map1["忍"] = 1
	for v, k := range map1 {
		fmt.Println(v, k)
	}

	fmt.Println(len(map1))
}

func sliceTest() {
	//// 生成切片
	//array := [5]string{"a", "b", "c", "d", "e"}
	//slice := array[2:5]

	//// 打印切片
	//fmt.Println(slice)
	//// 用索引访问切片元素:3.4.5 -> 0,1,2
	//fmt.Println(slice[1])
	//// 修改索引的值
	//slice[1] = "k"
	//fmt.Println(slice)
	//// 省略start 省略end
	//slice1 := array[:]
	//fmt.Println(slice1)

	// 切片声明
	//slice := make([]string, 4, 8)
	//fmt.Println(len(slice), cap(slice))
	slice1 := []string{"a", "b", "c", "z"}
	slice2 := append(slice1, "e", "f")
	slice3 := append(slice2, slice1...)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)

	// 切片循环
	for i, v := range slice3 {
		fmt.Print(i, "->", v, ";")
	}
}

func arrayTest() {
	//// 定义一个数组
	//array := [5]string{"a", "b", "c", "d", "e"}

	//// 省略长度
	//array := [...]string{"a", "b", "c", "d", "e"}

	// 不省略长度，但是一次只初始化两个索引
	array := [5]string{3: "a", 1: "b"}

	// 遍历数组
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}

	for i, v := range array {
		fmt.Printf("index %d is %s \n", i, v)
	}
}
