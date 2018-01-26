package main

import "fmt"

func sum(a []int, c chan int)  {
	total := 0
	for _,value := range a {
		total += value
	}
	c <- total // send total to c
}

func main() {
	a := []int{7,2,8,-9,4,0}
	c := make(chan int)
	go sum(a[:len(a)/2],c)
	go sum(a[len(a)/2:],c)

	x, y := <-c, <-c //receive from c
	fmt.Println(x,y,x+y)

	ch := make(chan int, 3) //指定channel的缓冲大小
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	ch <- 3
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

