package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v := <-c:
				fmt.Println(v)
			case <- time.After(time.Second*5):
				fmt.Println("timeout")
				o <- true
			}
		}
	}()
	<- o
}

// Goexit 退出当前执行的Goroutine。但是defer函数还是会继续调用
// Gosched 让出当前Goroutine的执行权限，调度器安排其他等待的任务运行。并在下次某个时候从该位置恢复执行
// NumCPU 返回CPU核心数
// NumGoroutine 返回正在执行和排队的任务总数
// GOMAXPROCS 用来设置可以并行计算的CPU核数的最大值，并返回之前的值。