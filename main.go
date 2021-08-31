package main

import (
	"log"
	"net/http"
	"text/template"
)

var plantillas = template.Must(template.ParseGlob("plantillas/*"))

func main() {
	http.HandleFunc("/", Inicio)
	log.Println("Servidor corriendo.")
	http.ListenAndServe(":8081", nil)
}

func Inicio(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hola mundo.")
	plantillas.ExecuteTemplate(w, "inicio", nil)
}
