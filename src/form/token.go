package main

import (
	"net/http"
	"fmt"
	"time"
	"crypto/md5"
	"io"
	"strconv"
	"html/template"
)

func index(w http.ResponseWriter, r *http.Request)  {
	r.PostForm()
	fmt.Fprintf(w,"Hello golang!")
}

func login(w http.ResponseWriter, r *http.Request)  {
	if r.Method == "GET" {
		curTime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(curTime,10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t,_ := template.ParseFiles("src/form/token.tmpl")
		t.Execute(w, token)
	}else {
		// 请求的是登录数据，执行登录的逻辑判断
		r.ParseForm()
		token := r.Form.Get("token")
		if token != ""{
			// 验证token合法性
		}else{
			// 不存在token报错
		}
		fmt.Println("username length",len(r.Form["username"][0]))

		//输出到服务器端
		fmt.Println("username: ",template.HTMLEscapeString(r.Form.Get("username")))
		fmt.Println("password: ",template.HTMLEscapeString(r.Form.Get("password")))
		//输出到客户端
		template.HTMLEscape(w, []byte(r.Form.Get("username")))
	}
}

func main() {
}
