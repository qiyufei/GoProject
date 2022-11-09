package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//testGo()
	//testChannel()
	testSelect()
}

func testSelect() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)
	go func() {
		file := downloadFile("r1")
		ch1 <- file
	}()
	go func() {
		file := downloadFile("r2")
		ch2 <- file
	}()
	go func() {
		file := downloadFile("r3")
		ch3 <- file
	}()
	select {
	case file := <-ch1:
		fmt.Println(file)
	case file := <-ch2:
		fmt.Println(file)
	case file := <-ch3:
		fmt.Println(file)
	}
}
func downloadFile(routineName string) string {
	time.Sleep(time.Duration(rand.Uint64()))
	return routineName
}

func testChannel() {
	// 创建一个channel, 两个goroutine 通信
	ch := make(chan string)
	go func() {
		fmt.Println("yufei")
		ch <- "来自 yufei routine"
	}()
	fmt.Println("ch22 routine")
	s := <-ch
	fmt.Println(s)
	close(ch)
}

func testGo() {
	go fmt.Println("yufei")
	fmt.Println("ch22 goroutine")
	time.Sleep(1000)
}
