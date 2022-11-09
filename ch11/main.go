package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//selectTimeout()
	//testPipeline()
	//testMerge()
	testFutures()
}

//搞两个协程两个管道，我自己该干啥干啥
func testFutures() {
	vegetablesCh := make(chan string)
	go func() {
		time.Sleep(5 * time.Second)
		vegetablesCh <- "wash vegetables"
	}()
	waterCh := make(chan string)
	go func() {
		time.Sleep(5 * time.Second)
		waterCh <- "boil water"
	}()
	fmt.Println("已经安排洗菜和烧水了，我先眯一会")
	time.Sleep(2 * time.Second)
	fmt.Println("要做火锅了，看看菜和水好了吗")
	vegetables := <-vegetablesCh
	water := <-waterCh
	fmt.Println("准备好了，可以做火锅了:", vegetables, water)
}

func testMerge() {
	coms := buy(100)
	phone1 := build(coms)
	phone2 := build(coms)
	phone3 := build(coms)

	phones := merge(phone1, phone2, phone3)
	out := pack(phones)
	for o := range out {
		fmt.Println(o)
	}
}

// 多个channel中的数据发送到一个channel中
func merge(ins ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	wg.Add(len(ins))
	out := make(chan string)

	for _, in := range ins {
		go func(in <-chan string) {
			defer wg.Done()
			for c := range in {
				out <- c
			}
		}(in)
	}

	go func() {
		defer close(out)
		wg.Wait()
	}()
	return out
}

// 搭建流水线：购买-> 组装 -> 包装
func testPipeline() {
	for c := range pack(build(buy(10))) {
		fmt.Println(c)
	}
}
func buy(num int) <-chan string { //只能被读取的管道，不能写
	out := make(chan string)
	go func() {
		defer close(out)
		for i := 0; i < num; i++ {
			out <- fmt.Sprint("配件", i)
			time.Sleep(time.Second)
		}
	}()
	return out
}
func build(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		// 读取channel，直到channel关闭
		for c := range in {
			out <- "组装（" + c + ")"
		}
	}()
	return out
}
func pack(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for c := range in {
			out <- "pack (" + c + ")"
		}
	}()
	return out
}

// test timeout
func selectTimeout() {
	ch := make(chan string)
	go func() {
		time.Sleep(8 * time.Second)
		ch <- "ok"
	}()
	select {
	case v := <-ch:
		fmt.Println(v)
	case <-time.After(5 * time.Second):
		fmt.Println("time out")
	}
}
