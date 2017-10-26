package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	width, height float64
}

type Circleangle struct {
	radius float64
}

func(r Rectangle) area()float64  {
	return r.width * r.height
}

func(c Circleangle) area()float64  {
	return c.radius * c.radius * math.Pi
}


func main() {
	r1 := Rectangle{12,2}
	r2 :=Rectangle{3,6}
	c1 := Circleangle{6}
	c2 := Circleangle{9}
	fmt.Println("Area of r1 is ",r1.area())
	fmt.Println("Area of r2 is ",r2.area())
	fmt.Println("Area of c1 is ",c1.area())
	fmt.Println("Area of c2 is ",c2.area())
}
