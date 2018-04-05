package main

import (
	"html/template"
	"net/http"
	"time"
	"log"
	// "controller"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}


func main() { 

	mux := http.NewServeMux()

	mux.HandleFunc("/", index)
	mux.HandleFunc("/agregar_vecino", agregar_vecino)
	mux.HandleFunc("/visita", visita)

	server := &http.Server{
		Addr:				":8080",
		Handler:			mux,
		ReadTimeout:		10 * time.Second,
		WriteTimeout: 		10 * time.Second,
		MaxHeaderBytes: 	1<<20,
	}

	log.Println("escuchando...")

	log.Fatal(server.ListenAndServe())
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func agregar_vecino(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "agregar_vecino.gohtml", nil)
}

func visita(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "visita.gohtml", nil)
}