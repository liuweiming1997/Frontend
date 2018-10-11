package main

import (
	"fmt"
	"log"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的

	fmt.Println("get message from " + r.RemoteAddr)
	fmt.Println(r.Form)
	fmt.Fprintf(w, "I get your message!") //这个写入到w的是输出到客户端的
}

func main() {
	http.HandleFunc("/", root)               //设置访问的路由
	err := http.ListenAndServe(":8080", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
