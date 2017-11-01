package main

import (
	"io/ioutil"
	"fmt"
	"os"
)

func main() {
	inputFile := "src/file/input.txt"
	outputCopyFile := "src/file/input_copy.txt"
	buf,err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr,"File error: %s\n",err)
	}
	fmt.Printf("%s\n",string(buf))
	err = ioutil.WriteFile(outputCopyFile, buf, 0644)
	if err != nil {
		panic(err.Error())
	}
}
