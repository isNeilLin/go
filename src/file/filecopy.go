package main

import (
	"os"
	"io"
	"fmt"
)

func main() {
	CopyFile("src/file/fileinput_copy.txt","src/file/fileinput.go")
	fmt.Println("copy done!")
}

func CopyFile(target, source string)(written int64,err error)  {
	src, err := os.Open(source)
	if err != nil {
		panic(err.Error())
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(target,os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err.Error())
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}