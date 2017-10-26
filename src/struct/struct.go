package main

import "fmt"


//声明一个新的类型
type person struct{
	name string
	age int
}

type skills []string // 自定义类型skills

type student struct {
	person		// 匿名字段，默认student包含了person的所有字段
	socre int
}

type worker struct {
	student
	skills
	int // 内置类型作为匿名字段
}
//所有的内置类型和自定义类型都可以作为匿名字段


func older(p1,p2 person)(person, int)  {
	if p1.age > p2.age{
		return p1,p1.age - p2.age
	}
	return  p2,p2.age - p1.age
}

func main() {
	var tom person
	tom.name,tom.age = "tom",18

	bob := person{age: 20, name: "bob"}

	paul := person{"paul",25}

	stu1 := student{tom,80}

	engineer := worker{stu1,skills{"node","html","css","go","python"},int(12)}

	var skill string
	for _,value := range engineer.skills {
		skill += value+","
	}

	tb_older, tb_diff := older(tom,bob)
	tp_older, tp_diff := older(tom,paul)
	bp_older, bp_diff := older(bob,paul)

	fmt.Printf("%s and %s, %s is older by %d years\n",
		tom.name, bob.name, tb_older.name, tb_diff)
	fmt.Printf("%s and %s, %s is older by %d years\n",
		tom.name, paul.name, tp_older.name, tp_diff)
	fmt.Printf("%s and %s, %s is older by %d years\n",
		paul.name, bob.name, bp_older.name, bp_diff)
	fmt.Printf("%s is a student,%d last week\n",stu1.name,stu1.socre)
	fmt.Printf("%s is a engineer,%s's skills is %s and got %d each month\n",
		engineer.name,engineer.name,skill,engineer.int)
}
