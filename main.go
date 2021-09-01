package main

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"

	_ "github.com/mattn/go-sqlite3"
)

func conexionBD() (conexion *sql.DB) {
	Driver := "sqlite3"
	database := "db/hoard.db"

	conexion, err := sql.Open(Driver, database)
	if err != nil {
		panic(err.Error())
	}
	return conexion
}

var plantillas = template.Must(template.ParseGlob("plantillas/*"))

func main() {
	http.HandleFunc("/", Inicio)
	http.HandleFunc("/actor", Actor)

	log.Println("Servidor corriendo.")
	http.ListenAndServe(":8081", nil)
}

func Inicio(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hola mundo.")
	conexionEstablecida := conexionBD()
	insertarRegistros, err := conexionEstablecida.Prepare("INSERT INTO actores(nombre,edad) VALUES ('John',34)")
	if err != nil {
		panic(err.Error())
	}
	insertarRegistros.Exec()

	plantillas.ExecuteTemplate(w, "inicio", nil)
}

func Actor(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hola mundo.")
	plantillas.ExecuteTemplate(w, "crear", nil)
}
