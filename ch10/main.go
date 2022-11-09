package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 启动一个监控狗
// 通过channel让监控狗停下
func main() {
	//channelStopRoutine()
	//contextStopRoutine()
	//stopContextTree()
	testValueContext()
}
func testValueContext() {
	var wg sync.WaitGroup
	wg.Add(1)
	valCtx := context.WithValue(context.Background(), "username", "yufei")
	go func() {
		defer wg.Done()
		getUser(valCtx)
	}()

	time.Sleep(5 * time.Second)
	// 还要通过结束ctx，才能走到return分支。
	wg.Wait() // 如果不执行wg.done 就会一直wait。如果不return getUser就会一直不执行defer。
}

func getUser(ctx context.Context) {
	a := ctx.Value("username")
	for {
		select {
		case <-ctx.Done():
			fmt.Println("停止")
			return
		default:
			fmt.Println("username", a)
			fmt.Println("在监控")
		}
		time.Sleep(time.Second)
	}
}

func stopContextTree() {
	var wg sync.WaitGroup
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())

	go func() { // 开启一个协程
		defer wg.Done()
		watchDog3(ctx, "协程1")
	}()

	time.Sleep(5 * time.Second)
	cancel()
	wg.Wait()
}

func watchDog3(ctx context.Context, name string) {
	// 再去开启一个协程，然后使用 context 去控制取消。使用ctx 再去生成一个context，
	ctx2 := context.Background()
	go func() {
		defer fmt.Println("协程1开的协程2", "结束") //这里不会执行，ctx2的done也不会执行
		watchDog2(ctx2, "协程1开的协程2")
	}()
	for {
		select {
		case <-ctx.Done(): //要有<-
			fmt.Println("收到停止命令")
			return
		default:
			fmt.Println(name, "正在监控...")
		}
		time.Sleep(1 * time.Second)
	}
}

func contextStopRoutine() {
	var wg sync.WaitGroup
	wg.Add(3)
	ctx, cancelFunc := context.WithCancel(context.Background())

	go func() {
		defer wg.Done()
		watchDog2(ctx, "watch dog 1")
	}()

	// start a watch dog
	go func() {
		defer wg.Done()
		watchDog2(ctx, "监控狗2")
	}()

	go func() {
		defer wg.Done()
		watchDog2(ctx, "watch dog 3")
	}()

	time.Sleep(5 * time.Second)
	cancelFunc() // 上面返回的cancelFunc
	wg.Wait()
}
func watchDog2(ctx context.Context, name string) { //类型是context.Context
	for {
		select {
		case <-ctx.Done(): //要有<-
			fmt.Println(name, "收到停止命令")
			return
		default:
			fmt.Println(name, "正在监控...")
		}
		time.Sleep(1 * time.Second)
	}
}

func channelStopRoutine() {
	// 创建一个waitGroup让main routine不结束
	wg := sync.WaitGroup{}
	wg.Add(1)
	finishChan := make(chan bool)

	go func() {
		defer wg.Done()
		watchDog1("监控狗1", finishChan)
	}()

	time.Sleep(5 * time.Second)
	finishChan <- true
	wg.Wait()
}
func watchDog1(name string, finishChan chan bool) {
	for {
		select {
		case <-finishChan:
			fmt.Println("收到停止命令")
			return
		default:
			fmt.Println(name, "正在监控...")
		}
		time.Sleep(1 * time.Second)
	}
}
