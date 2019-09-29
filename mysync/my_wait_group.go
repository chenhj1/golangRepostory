package main

import (
	"fmt"
	"sync"
	"time"
)

// ---------------------------------------------------------------------------------------------------------------------
// waitGroup 用于等待一组 goroutine 结束。
// 主协程调用wg.wait()方法之后，阻塞在此处直至 waitGroup 中的groutine执行完毕，然后主程序继续执行
//
// note: wg.Add()需要在goutine开始执行之前调用。特别注意不要在goutine开始执行之后调用。
// note: wg作为函数参数时，必须是指针类型，否则失效！
// ---------------------------------------------------------------------------------------------------------------------

func funwg(i int, wg *sync.WaitGroup) {
	fmt.Println("goroutine:", i)
	time.Sleep(3 * time.Second)
	wg.Done()
}

func main() {
	wg := &sync.WaitGroup{}

	for i := 0; i < 3; i++ {
		wg.Add(1) // 特别注意：wg.Add() 需要在goroutine开始执行之前调用。
		go funwg(i, wg)
	}

	wg.Wait()
	fmt.Println("all done.")
}
