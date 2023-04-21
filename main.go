package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"ascii-art/function"
)

func Errors(w http.ResponseWriter, code int, message string) {
	t, err := template.ParseFiles("templates/error.html")
	if err != nil {
		log.Println("Error parsing template")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(code)
	response := struct {
		ErrorCode int
		ErrorText string
	}{
		ErrorCode: code,
		ErrorText: message,
	}
	t.Execute(w, response)
}

func ascii_Art(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" {
		Errors(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}
	switch r.Method {
	case http.MethodPost:
		asciiPost(w, r)
	default:
		Errors(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		Errors(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}
	switch r.Method {
	case http.MethodGet:
		homeGet(w, r)
	default:
		Errors(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
	}
}

func homeGet(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Println("Error parsing template at")
		Errors(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Println("Error executing template")
		Errors(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
}

func asciiPost(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Println("Error parsing template")
		Errors(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	word := r.FormValue("check")
	if r.FormValue("font") != "shadow" && r.FormValue("font") != "thinkertoy" && r.FormValue("font") != "standard" {
		Errors(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	if len(word) > 400 {
		Errors(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	font := "banners/" + r.FormValue("font") + ".txt"
	if !(r.Form.Has("font") && r.Form.Has("check")) {
		Errors(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	standFont, err := function.ReadFont(font)
	if err != nil {
		log.Println(err.Error())
		Errors(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	if !function.CheckHash(function.MD5(font), font) {
		log.Println("Banner has been changed")
		Errors(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	answer, err := function.PrintFormat(word, standFont)
	if err != nil {
		Errors(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	err = t.Execute(w, answer)
	if err != nil {
		log.Println("Error executing template")
		Errors(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
}

func main() {
	args := os.Args[1:]
	if len(args) > 0 {
		fmt.Println("Usage: go run .")
		return
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/ascii-art", ascii_Art)
	mux.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("static"))))
	log.Println("Localhost on port 8080:")

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	log.Fatal(err)
}
