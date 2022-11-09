package main

import "fmt"

func main() {
	var s string
	s = "yufei"
	fmt.Println(s)
	fmt.Println(&s)

	var sp *string
	sp = new(string)
	fmt.Println(*sp, sp)
}
