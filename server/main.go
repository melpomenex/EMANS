package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type loginDetails struct {
	Username string
	Password string
}

func login(w http.ResponseWriter, req *http.Request) {
	u := req.FormValue("username")
	p := req.FormValue("password")

	err := tpl.ExecuteTemplate(w, "login.gohtml", loginDetails{u, p})
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}

func index(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.ListenAndServe(":8080", nil)
}
