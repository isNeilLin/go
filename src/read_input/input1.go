package main

import "fmt"

var (
	i int
	f float32
	s string
	input = "56.12 / 5212 /GO"
	format = "%f / %d /%s"
)
func main() {
	firstName,lastName := "",""
	fmt.Println("Please type your full name: ")
	fmt.Scanln(&firstName,&lastName)
	fmt.Printf("Hi %s %s!\n",firstName, lastName)
	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println(f,i,s)
}
