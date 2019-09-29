package main

import (
	"fmt"
	"sync"
	"time"
)

// sync包中的Wait和Signal使用
//
// Wait: 使当前goroutine阻塞，直到被其他的goroutine通知（即调用Singal方法）
// Signal: 唤醒一个wait()中的协程
// SignalAll：唤醒全部Wait()的协程
//
// 与java的wait和notify如出一辙:
// 相同点：java的wait和notify必须在synchronized代码块中使用；golang的wait和signal必须在mutex中使用。二者都需要加锁。（golang的signal不需要在锁中调用也可以，不强制）
// 不同点：java的wait和notify是对象的方法；golang的wait和signal需要定义一个 sync.Cond 变量，传入一个 Mutex 锁来使用。即：java是object.wait(); golang是cond.Wait()

func ft(i int, l *sync.Mutex, c *sync.Cond) {
	fmt.Println("Not lock", i)
	l.Lock()

	fmt.Println("Locked ", i)
	fmt.Println("before wait ", i)
	fmt.Println()
	time.Sleep(1 * time.Second)
	c.Wait()

	fmt.Println("after wait ", i)
	fmt.Println("Unlock ", i)
	l.Unlock()
}

func main() {

	l := &sync.Mutex{}   // 定义一把锁
	c := sync.NewCond(l) // cond变量用于控制golang控制wait和signal。传入一把锁作为参数。

	// 定义3个协程，竞争抢锁。抢到了锁，率先进入wait等待。
	for i := 0; i < 3; i++ {
		go ft(i, l, c)
	}

	// 循环3次，每次 signal “唤醒”一个协程。循环3次刚好逐个唤醒上面的所有协程
	for i := 0; i < 3; i++ {
		time.Sleep(5 * time.Second)
		c.Signal()
	}

	time.Sleep(time.Hour) // 防止main函数退出，无实际意义
}
