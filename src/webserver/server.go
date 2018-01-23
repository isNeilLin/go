package main

import (
	. "fmt"
	"net/http"
	"html/template"
	"log"
	"strconv"
	"regexp"
)

type any struct {
	msg string
	pass bool
}

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数
	Fprintln(w, "Hello Golang")
	Fprintln(w, "Hello World")
}

func isChinese(s string) bool {
	if m,_ := regexp.MatchString("^\\p{Han}+$",s); !m {
		return false
	} else {
		return true
	}
}

func volidator(r *http.Request) any {
	if len(r.Form["username"][0])==0 {
		return any{"用户名不能为空", false}
	}
	getint,err := strconv.Atoi(r.Form.Get("age"))
	if err != nil {
		return any{"无效年龄", false}
	}
	if getint > 100 {
		return any{"无效年龄", false}
	}
	return any{"验证通过", true}
}

func login(w http.ResponseWriter, r *http.Request) {
	Println("method: ", r.Method)
	r.ParseForm()
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		result := volidator(r)
		if result.pass {
			Println(r.Form["username"])
			Println(r.Form["password"])
		} else {
			Println(result.msg)
		}
	}
}

func main() {
	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/login",login)
	err := http.ListenAndServe(":7000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	Println("Server listened at port localhost: 7000")
}