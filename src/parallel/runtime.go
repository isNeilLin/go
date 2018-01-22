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

func main() {
	go say("Golang") //开一个新的Goroutine执行
	go speak("World")
	say("hello") //当前Goroutine执行
}
