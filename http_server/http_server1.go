package main

import (
	"net/http"
	"fmt"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
	})

	// Serving static assets(资源)
	// 这个地方简单的说就是文件服务器
	// 后面这个 路径应该怎么写
	// 绝对路径肯定是可以的,问题是 相对路径如何写?  这个相对路径是对 可执行 main 的程序来的，而不是当前在 IDE中运行的临时的
	path := "static/"
	//path = "/Users/penglonghua/Desktop"
	fs := http.FileServer(http.Dir(path))  // 这个地方的返回值就是一个 handler


	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":9090", nil)


}
