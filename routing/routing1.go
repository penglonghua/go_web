package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
)

func main() {


	// 1 新建 router

	r := mux.NewRouter()

	// 2 注册
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {

		// 这个地方可以从 http.Request中获取
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	//  Features
	// 1严格限制方法
	r.HandleFunc("/books/{title}/post", CreateBook).Methods("POST")
	r.HandleFunc("/books/{title}/read", CreateBook).Methods("GET")

	// 2 严格限制域名或者 hostname 才能访问, 这个是什么意思?
	r.HandleFunc("/books/mybookstore", BookHandler).Host("www.mybookstore.com")  // 没有跑通

	// 3 限定 request to http/https
	r.HandleFunc("/secure", SecureHandler).Schemes("https")  //  没有跑出来
	r.HandleFunc("/insecure", InsecureHandler).Schemes("http")

	// 4 path prefixex & Subrouters
	bookrounter := r.PathPrefix("/pens").Subrouter()
	bookrounter.HandleFunc("/", AllPens)
	bookrounter.HandleFunc("/{title}", GetPen)

	http.ListenAndServe(":9090", r)




}



func CreateBook(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "方法 %s\n", r.Method)
}

func BookHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "方法 %s\n", r.Method)
}

func SecureHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "https")
}


func InsecureHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "http")
}


func AllPens(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "所有的笔")
}

func GetPen(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "某一支笔")
}