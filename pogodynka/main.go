package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("uruchomiono serwer")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	dane, _ := weather()
	t, _ := template.ParseFiles("weather_template.html")
	t.Execute(w, dane)
}
