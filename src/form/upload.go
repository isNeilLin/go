package main

import (
	"net/http"
	"time"
	"fmt"
	"crypto/md5"
	"io"
	"strconv"
	"html/template"
	"log"
	"os"
)

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		curTime := time.Now().Unix()
		fmt.Println(curTime)
		h := md5.New()
		fmt.Println(h)
		io.WriteString(h,strconv.FormatInt(curTime,10))
		token := fmt.Sprintf("%x",h.Sum(nil))
		fmt.Println(token)
		t,_ := template.ParseFiles("src/form/upload.tmpl")
		t.Execute(w,token)
	}else{
		r.ParseMultipartForm(32<<20)
		file, handler, err := r.FormFile("uploadfile")
		if err!=nil {
			log.Fatal(err)
		}
		defer file.Close()
		fmt.Fprintf(w,"%v",handler.Header)
		f, err := os.OpenFile("src/form/test/"+handler.Filename,os.O_WRONLY|os.O_CREATE,0666)
		if err!=nil {
			log.Fatal(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}
func main() {
	http.HandleFunc("/upload",upload)
	err := http.ListenAndServe(":9090",nil)
	if err!=nil {
		log.Fatal(err)
	}
}
