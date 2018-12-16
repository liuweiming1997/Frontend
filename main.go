package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	rootPage = `
  <!doctype html>
  <html>
  <body>
  <form action='/saveImage' method='post' enctype='multipart/form-data'>
     <input type='file' name='vimi_image'>
     <input type='submit' value='Upload'>
  </form>
  `
	savePath = `./img/`
)

func inWhereAndThenShowCookies(name string, r *http.Request) {
	fmt.Println(name)
	for k, v := range r.Cookies() {
		fmt.Println(k, v)
	}
	fmt.Println("---->\n")
}

func root(w http.ResponseWriter, r *http.Request) {
	inWhereAndThenShowCookies("in root", r)
	r.ParseForm()            //解析参数，默认是不会解析的
	fmt.Fprintf(w, rootPage) //这个写入到w的是输出到客户端的
}

func saveImage(w http.ResponseWriter, r *http.Request) {
	inWhereAndThenShowCookies("in saveImage", r)
	fmt.Println("in saveImage")
	r.ParseForm()
	uploadFile, handle, err := r.FormFile("vimi_image")
	if err != nil {
		log.Println(err)
		return
	}
	os.Mkdir(savePath, 0777)
	saveFile, err := os.OpenFile(savePath+handle.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println(err)
		return
	}
	io.Copy(saveFile, uploadFile)
	defer uploadFile.Close()
	defer saveFile.Close()
	r.AddCookie(&http.Cookie{
		Name:    "fileName",
		Value:   handle.Filename,
		Expires: time.Now().Add(24 * time.Hour),
	})
	// must set it to response, and use in next request
	http.SetCookie(w, &http.Cookie{
		Name:    "fileName",
		Value:   handle.Filename,
		Expires: time.Now().Add(24 * time.Second),
	})
	http.Redirect(w, r, "/showImage", http.StatusFound)
}

func showImage(w http.ResponseWriter, r *http.Request) {
	inWhereAndThenShowCookies("in showImage", r)
	r.ParseForm()
	fileName, err := r.Cookie("fileName")
	if err != nil {
		log.Println(err)
		return
	}
	file, err := os.Open(savePath + fileName.Value)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	buff, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
		return
	}
	w.Write(buff)
}

func main() {
	http.HandleFunc("/", root)               //设置访问的路由
	http.HandleFunc("/saveImage", saveImage) //设置访问的路由
	http.HandleFunc("/showImage", showImage) //设置访问的路由
	err := http.ListenAndServe(":8080", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
