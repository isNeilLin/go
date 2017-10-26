package main

import "fmt"

type Human struct {
	name string
	age int
	phone string
}

type Student struct {
	Human
	school string
	loan float32
}

type Employee struct {
	Human
	company string
	money float32
}

func (h Human)SayHi()  {
	fmt.Printf("Hi, I am %s, you can call me on %s. \n",h.name, h.phone)
}

func (h Human)Sing(lrc string)  {
	fmt.Println("la la la la la ",lrc)
}

func (e Employee)SayHi()  {
	fmt.Printf("Hi, I am %s, I am work at %s, call me at %s.\n",
	e.name, e.company, e.phone)
}

type Men interface {
	SayHi()
	Sing(lrc string)
}

func main() {
	mike := Student{Human{"Mike",16,"860-221-4136"},"MIT",0.00}
	paul := Student{Human{"Paul",18,"832-155-5421"},"Harvard",100}
	sam := Employee{Human{"Sam",25,"943-123-4124"},"GoLang Inc",20000}
	tom := Employee{Human{"Tom",28,"823-535-4321"},"Apple Inc",50000}

	// 定义Men类型的变量
	var i Men

	// i能存储Student
	i = mike
	i.SayHi()
	i.Sing("I know I am gonna miss you I know")

	// i也能存储Employee
	i = sam
	i.SayHi()
	i.Sing("so, if you love someone, you should let then know")

	x := make([]Men, 4)

	x[0], x[1], x[2], x[3] = paul, sam, mike, tom

	for _,value := range x{
		value.SayHi()
	}
}
