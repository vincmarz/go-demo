package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Message struct {
	Message string
}

func index(w http.ResponseWriter, r *http.Request) {

	HomePageVars := Message{
		Message: os.Getenv("MESSAGE"),
	}

	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	err = t.Execute(w, HomePageVars)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}
func main() {

	http.HandleFunc("/", index)
	fmt.Println("Server starting...")
	http.ListenAndServe(":8080", nil)
}
