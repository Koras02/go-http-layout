package main

import (
	"html/template"
	"net/http"
)

type PageVariables struct {
	Title string
}

func main() {
	// 정적 파일을 사용하기 위한 핸들러 설정
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	
	http.HandleFunc("/", HomePage)
	http.ListenAndServe(":8080", nil);
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	PageVariables := PageVariables{
		Title: "My WebSite",
	}
	tmpl, err := template.ParseFiles("templates/layout.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}
	tmpl.Execute(w, PageVariables)
}