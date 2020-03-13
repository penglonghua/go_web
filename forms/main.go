package main

import (
	"html/template"
	"net/http"
)

type ContactDetails struct {

	Email string
	Subject string
	Message string

}

func main() {

	tmpl := template.Must(template.ParseFiles("forms.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// 判断一下是否是 Post方法
		if r.Method != http.MethodPost{
			tmpl.Execute(w, nil)  //空数据也可以的
			return
		}

		details := ContactDetails{
			Email: r.FormValue("email"),
			Subject : r.FormValue("subject"),
			Message : r.FormValue("message"),

		}

		// do something with details
		_ = details

		tmpl.Execute(w, struct {Success bool}{true})

	})

	http.ListenAndServe(":9090", nil)


}