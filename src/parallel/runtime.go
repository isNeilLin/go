package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1*time.Millisecond) //表示让CPU把时间片让给别人，下次某个时候继续恢复执行该goroutine
		fmt.Println(s)
	}
}

func speak(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1*time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)
	go say("Golang") //开一个新的Goroutine执行
	go speak("World")
	go sum(s[:len(s)/2], c1)
	go sum(s[len(s)/2:], c2)
	go sum(s, c3)
	say("hello") //当前Goroutine执行
	x, y, z := <-c1, <-c2, <-c3 // receive from c
	fmt.Println(x,y,z)
}
