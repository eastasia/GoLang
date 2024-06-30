package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {

	http.HandleFunc("/lookup", formHandler)

	dir := http.Dir("C:/Users/mike_/go/src/server/htdocs")
	fileServer := http.FileServer(dir)
	http.Handle("/", fileServer)

	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting web server on port 8000")
	http.ListenAndServe(":8000", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Response from the helloHandler!\n")
	io.WriteString(w, "URL.Path = "+r.URL.Path)
}

func formHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", 405)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", 400)
		return
	}

	lang := r.FormValue("Language")
	lang = strings.ToLower(strings.Trim(lang, " "))

	file, err := os.OpenFile("htdocs/data.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Print(err)
	}
	defer file.Close()
	log.SetOutput(file)
	log.Print(lang)

	title := "<a href='http://www.ineasysteps.com'>See all titles online</a>"
	isbn := "logo"

	switch lang {
	case "c":
		title = "C Programming"
		isbn = "978-1-84078-840-2"
	case "go":
		title = "Go Programming"
		isbn = "978-1-84078-919-5"
	case "java":
		title = "Java"
		isbn = "978-1-84078-873-0"
	case "python":
		title = "Python"
		isbn = "978-1-84078-812-9"
	case "sql":
		title = "SQL"
		isbn = "978-1-84078-902-7"
	}

	response := "<!DOCTYPE HTML><title>Web Server Response</title>"
	response += "<link rel='shortcut icon' href='favicon.ico'>"
	response += "<link rel='stylesheet' href='css/style.css'>"
	response += "<p>" + title + "<br>in easy steps"
	if isbn != "logo" {
		response += "<br>ISBN:" + isbn
	}
	response += "</p><img src='img/" + isbn + ".png' alt='Cover'>"
	response += "<img src='img/power.png' alt='Power'>"
	io.WriteString(w, response)
}
