package main

import (
	"html/template"
	"net/http"
)

type Name struct {
	FName []string
	Color string
}

func main() {
	http.HandleFunc("/", helloWorld)
	_ = http.ListenAndServe(":8888", nil)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	name := Name{
		FName: []string{"Bishwajit", "Open", "Source"},
		Color: "limegreen",
	}
	temp, _ := template.ParseFiles("hello.gohtml")
	_ = temp.Execute(w, name)
}
