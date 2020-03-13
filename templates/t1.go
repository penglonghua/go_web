package main

import (
	"html/template"
	"net/http"
)

type Todo struct {
	Title string
	Done bool
}

type TodoPageData struct {
	PageTitle string
	Todos []Todo
}

func main() {

	// parse template from files
	// tmpl, err := template.ParseFiles("layout.html")
	tmpl := template.Must(template.ParseFiles("layout.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos :[]Todo{
				{Title: "Task 1", Done: false },
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},

		}
		// 这个地方 自然会有  Content-Type: text/html; charset=utf-8.
		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":9090", nil)


}

