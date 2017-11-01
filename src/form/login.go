package main

import (
	"net/http"
	"fmt"
	"strings"
	"html/template"
	"path/filepath"
	"os"
	"log"
)

func index(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm() //解析url传递的参数。对于Post则解析x响应包的主体
	// 如果没有调用ParseForm方法，下面无法获取表单的数据
	for k,v := range r.Form {
		fmt.Println("key: ",k)
		fmt.Println("value: ",strings.Join(v,""))
	}
	fmt.Fprintf(w,"Hello astaxie!")
}

func login(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	if r.Method == "GET" {
		cwd,_ := os.Getwd()
		path := filepath.Join(cwd, "src/form/login.tmpl")
		t,err := template.ParseFiles(path)
		if err != nil {
			log.Fatal(err)
			return
		}
		t.Execute(w, nil)
	}else {
		fmt.Println("username",r.Form["username"])
		fmt.Println("password",r.Form["password"])
	}
}

func main() {
	http.HandleFunc("/",index)
	http.HandleFunc("/login",login)
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe:",err)
	}
}
