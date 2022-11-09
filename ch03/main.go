package main

import "fmt"

func main() {
	//if i := 5; i < 5 {
	//	fmt.Println("i < 10")
	//} else if i >= 5 && i < 10 {
	//	fmt.Println("i >= 5 && i < 10")
	//} else {
	//	fmt.Println("i>=10")
	//}

	//switch i := 6; {
	//case i < 5:
	//	fmt.Println("<5")
	//case i < 10:
	//	fmt.Println("< 10")
	//default:
	//	fmt.Println(i)
	//}

	//switch i := 1; i {
	//case 1:
	//	fallthrough
	//case 2:
	//	fmt.Println("符合条件")
	//default:
	//	fmt.Println("没有命中")
	//}
	//
	//switch 1 < 2 {
	//case true:
	//	fmt.Println("1 < 2")
	//case false:
	//	fmt.Println("1 >= 2")
	//default:
	//}
	//
	//switch {
	//case 1 >= 2:
	//	fmt.Println("1 >= 2")
	//case 1 < 2:
	//	fmt.Println("1 < 2")
	//default:
	//}

	//sum := 0
	//for i := 1; i <= 100; i++ {
	//	sum += i
	//}

	//i := 1
	//sum := 0
	//for i <= 100 {
	//	sum += i
	//	i++
	//}

	//i := 1
	//sum := 0
	//for {
	//	sum += i
	//	i++
	//	if i > 100 {
	//		break
	//	}
	//}
	//fmt.Println(sum)

	i := 0
	sum := 0
	for i <= 100 {
		i++
		if i%2 != 0 {
			continue
		}
		sum += i
	}
	fmt.Println(sum)
}
