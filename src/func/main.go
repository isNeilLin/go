package main

import "fmt"

type testInt func(int) bool //声明了一个函数类型

func isOdd(X int) bool  {
	if X % 2 == 0{
		return false
	}
	return true
}

func isEven(X int) bool  {
	if X % 2 == 0{
		return true
	}
	return false
}

//声明的函数类型在这个地方当作了一个参数
func filter(slice []int, f testInt)[]int  {
	var result []int
	for _,value := range slice {
		if f(value){
			result = append(result,value)
		}
	}
	return result
}

func main()  {
	slice := []int {1,2,3,4,5,7}
	fmt.Println("slice = ",slice)
	odd := filter(slice,isOdd)
	fmt.Println("odd element = ",odd)
	even := filter(slice,isEven)
	fmt.Println("even element = ",even)
}
