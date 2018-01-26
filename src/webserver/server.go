package main

import (
	"os"
	"io"
	"crypto/md5"
	"time"
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
	if r.Method == "GET" {
		curtime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h,strconv.FormatInt(curtime,10))
		token := Sprintf("%x",h.Sum(nil))
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
			//验证token的合法性
		} else {
			//不存在token报错
		}
		Println("username length:", len(r.Form["username"][0]))
		Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) //输出到服务器端
		Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username"))) //输出到客户端
	}
}

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		curtime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(curtime,10))
		token := Sprintf("%x",h.Sum(nil))

		t,_ := template.ParseFiles("upload.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20) // 上传的文件存储在maxMemory大小的内存里面，如果文件大小超过了maxMemory，那么剩下的部分将存储在系统的临时文件中
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			Println(err)
			return
		}
		defer file.Close()
		Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./test/"+handler.Filename,os.O_WRONLY|os.O_CREATE,0666)
		if err != nil {
			Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}

func main() {
	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/login",login)
	http.HandleFunc("/upload",upload)
	err := http.ListenAndServe(":7000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	Println("Server listened at port localhost: 7000")
}