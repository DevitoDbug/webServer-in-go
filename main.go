package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprintf(w, "POST request successful\n")
	fName := r.FormValue("fname")
	sName := r.FormValue("sname")

	_, err := fmt.Fprintf(w, "First name :%v \n", fName)
	if err != nil {
		log.Fatalf("%v\n", err)
		return
	}
	_, err = fmt.Fprintf(w, "Second Name: %v", sName)
	if err != nil {
		log.Fatalf("%v\n", err)
		return
	}

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "400 bad request", http.StatusBadRequest)
		return
	}

	w.Header().Add("content-type", "text/html")
	_, err := fmt.Fprintf(w, "hello")
	if err != nil {
		log.Printf("%v \n", err)
		return
	}

}

func main() {

	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)

	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port: 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
