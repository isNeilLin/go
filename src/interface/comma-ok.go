package main

import (
	"fmt"
	"strconv"
)

type Element interface {
	
}

type List[] Element

type Person struct {
	name string
	age int
}

func (p Person)String()string  {
	return "(name:"+p.name+"-age:"+strconv.Itoa(p.age)+" years)"
}

func main() {
	list := make(List,3)
	list[0] = 1
	list[1] = "hello"
	list[2] = Person{"Neil",23}

	// comma-ok断言

	for index,element := range list{
		if value,ok := element.(int); ok{
			fmt.Printf("list[%d] is an int and its value is %d\n",index,value)
		}else if value,ok := element.(string); ok {
			fmt.Printf("list[%d] is a string and its value is %s\n",index,value)
		}else if value,ok := element.(Person); ok {
			fmt.Printf("list[%d] is a Person and its value is %s\n",index,value)
		}else{
			fmt.Printf("list[%d] is a different type\n",index)
		}
	}

	// switch语法实现
	for index,element := range list{
		switch value := element.(type){	//element.(type)语法不能在switch外的任何逻辑里使用
		case int:
			fmt.Printf("list[%d] is an int and its value is %d\n",index,value)
		case string:
			fmt.Printf("list[%d] is a string and its value is %s\n",index,value)
		case Person:
			fmt.Printf("list[%d] is a Person and its value is %s\n",index,value)
		default:
			fmt.Printf("list[%d] is a different type\n",index)
		}
	}
}
