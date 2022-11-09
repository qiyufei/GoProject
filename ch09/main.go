package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	sum    = 0
	mutex  sync.Mutex
	rMutex sync.RWMutex
	cond   = sync.NewCond(&sync.Mutex{})
)

func main() {
	//testParallel()
	//testSyncMutex()
	//testRWMutex()
	//testWaitGroup()
	//testSyncOnce()
	testSyncCond()
}
func testCond(num int) {
	fmt.Println(num, "号已经lock")
	cond.L.Lock()

	fmt.Println(num, "已经wait")
	cond.Wait() // wait在这里的话，锁就释放了，唤醒的时候，会重新获取锁。wait的时候锁释放，然后其他协程就获取锁了。

	cond.L.Unlock()
	fmt.Println(num, "已经unlock")
}
func testSyncCond() {
	// 创建一个 sync.Cond

	var wg sync.WaitGroup
	wg.Add(11)

	// 10个人wait
	for i := 0; i < 10; i++ {
		go func(num int) {
			defer wg.Done()
			testCond(num)
		}(i)
	}
	// 等待所有goroutine都进入wait状态
	time.Sleep(20 * time.Second)

	go func() {
		defer wg.Done()
		fmt.Println("开跑")
		cond.Broadcast()
	}()

	// 防止提前退出
	wg.Wait()
}

func testSyncOnce() {
	var once sync.Once
	ch := make(chan bool)
	onceFunc := func() {
		fmt.Println("只执行一次")
	}
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceFunc)
			ch <- true // 十个协程往管道放十次
		}()
	}

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch) // 从管道读十次
	}
}

func testWaitGroup() {
	var wg sync.WaitGroup // 声明就行，没赋值
	wg.Add(110)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			fmt.Println("routine")
		}()
	}
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			fmt.Println("read")
		}()
	}
	wg.Wait()
}

func testRWMutex() {
	// 读写锁，读与读之间不锁
	// 写并发互斥,保证最后是1000
	for i := 0; i < 100; i++ {
		go add3(10)
	}
	// 读不互斥，读的时候可以读，不可以写，提高了性能。
	// 但是并不能保证全写完才去读的。所以这里打印的不一定是1000
	for i := 0; i < 10; i++ {
		go fmt.Println(readNum())
	}
	time.Sleep(2 * time.Second)
	fmt.Println(readNum())
}
func readNum() int {
	rMutex.RLock()
	defer rMutex.RUnlock()
	re := sum
	return re
}
func add3(i int) {
	rMutex.Lock()
	defer rMutex.Unlock()
	sum += i
}

func testSyncMutex() {
	for i := 0; i < 100; i++ {
		go add2(10)
	}
	time.Sleep(time.Second)
	fmt.Println(sum)
}
func add2(i int) {
	mutex.Lock()
	defer mutex.Unlock()
	sum += i
}

func testParallel() {
	for i := 0; i < 100; i++ {
		go add1(10)
	}
	time.Sleep(time.Second)
	fmt.Println(sum)
}
func add1(i int) {
	sum += i
}
