package main

import (
	"net/http"
	"fmt"
	"strings"
	"log"
)

func sayhelloName(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm() // 解析参数，默认不会解析
	fmt.Println(r.Form)	//输出到服务器端的打印信息
	fmt.Println("path:",r.URL.Path)
	fmt.Println("scheme",r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k,v := range r.Form {
		fmt.Println("key: ",k)
		fmt.Println("val: ",strings.Join(v,""))
	}
	fmt.Fprintf(w, "Hello astaxie!")
}

func main() {
	http.HandleFunc("/",sayhelloName) //设置访问的路由
	err := http.ListenAndServe(":9090",nil)
	if err!=nil {
		log.Fatal("ListenAndServe: ",err)
	}
}
