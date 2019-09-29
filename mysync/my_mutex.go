package main

import (
	"fmt"
	"sync"
	"time"
)

// ---------------------------------------------------------------------------------------------------------------------
// mutex互斥锁。
//
// 使用 Lock() 加锁后，不能再继续对其加锁，直到利用 Unlock() 解锁后才能再加锁
// 在 Lock() 之前使用 Unlock() 会导致 panic 异常
//
// 在一个 goroutine 获得 Mutex 后，其他 goroutine 只能等到这个 goroutine 释放该 Mutex
// 已经锁定的 Mutex 并不与特定的 goroutine 相关联。这样可以利用一个 goroutine 对其加锁，再利用其他 goroutine 对其解锁
//
// note: mutex作为函数参数时，必须是指针类型，否则失效！
// ---------------------------------------------------------------------------------------------------------------------

// 加锁之后，休眠几秒，再解锁。
// note：整个休眠期间，其他 goroutine 是拿不到这个锁的，只能等待本 goroutine 解锁
func fc1(i int, l *sync.Mutex) {
	fmt.Println("Not lock ", i)
	l.Lock()
	fmt.Println("Locked ", i)
	time.Sleep(5 * time.Second)
	fmt.Println()

	fmt.Println("Unlock ", i)
	l.Unlock()
}

func main() {
	l := &sync.Mutex{}

	// 启动3个协程，抢夺锁。一旦被一个协程抢到锁，其他协程调用 Lock()方法只能等待，直到抢占到锁的协程解锁。
	for i := 0; i < 3; i++ {
		go fc1(i, l)
	}

	time.Sleep(time.Hour) // 防止main函数退出，无实际意义
}
