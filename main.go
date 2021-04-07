package main

import (
	"html/template"
	"log"
	"net/http"
)

type PageVariables struct {
	Date string
	Time string
}

func main() {
	http.HandleFunc("/", HomePage)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Print("template executing error: ", err) //log it
	}
}
