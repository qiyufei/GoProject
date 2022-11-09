package main

import "fmt"

func main() {
	age := 18
	modifyAge(&age)
	fmt.Println(age)
}
func modifyAge(age *int) {
	*age = 20
}

func testP() {
	name := "yufei"
	nameP := &name
	fmt.Println("name", name)
	fmt.Println("nameP", nameP)

	fmt.Println(*nameP)
	*nameP = "雨飞飞"
	fmt.Println(*nameP)
	fmt.Println(name)

	i := 12
	var intp *int
	intp = &i
	fmt.Println(intp)

	var ip *int = new(int)
	*ip = 10
	fmt.Println(ip)
}
