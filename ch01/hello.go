package main

import (
	"fmt"
	"strings"
)

func main() {
	//fmt.Println("hello 雨飞")

	//var i int = 10
	//fmt.Println(i)

	//var f32 float32 = 23.3
	//var f64 float64 = 2.3490859
	//fmt.Println("f32 is", f32, "f64 is", f64)

	//var bf bool = false
	//var bt bool = true
	//fmt.Println("bf is", bf, "bt is", bt)

	//var s1 string = "hello"
	//var s2 string = "雨飞"
	//fmt.Println("s1 + s2 is", s1+s2)
	//s1 += s2
	//fmt.Println(s1)

	//var j int32
	//var f1 float64
	//var bl bool
	//var s3 string //""
	//fmt.Println(j, f1, bl, s3)

	//k := 10
	//bf := false
	//ss := "hello"
	//fmt.Println(k, bf, ss)

	//i := 10
	//pi := &i
	//fmt.Println(*pi)
	//fmt.Println(pi)

	//const (
	//	one = iota
	//	two
	//	three
	//	four
	//)
	//fmt.Println(one, two, three, four)

	//i2s := strconv.Itoa(10)
	//s2i, err := strconv.Atoi(i2s)
	//fmt.Println(i2s, s2i, err)

	//f := 3.25
	//i2f := float64(10)
	//var f2i = int(f)
	//fmt.Println(i2f, f2i)

	s := "hello"
	fmt.Println(strings.HasPrefix(s, "h"))
	fmt.Println(strings.Index(s, "o"))
	fmt.Println(strings.ToUpper(s))
	fmt.Println(strings.Contains(s, "el"))
}
