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
	http.HandleFunc("/actor", Actores)
	http.HandleFunc("/insertar", Insertar)

	log.Println("Servidor corriendo.")
	http.ListenAndServe(":8081", nil)
}

type Actor struct {
	id     int
	nombre string
	edad   int
}

func Inicio(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hola mundo.")
	conexionEstablecida := conexionBD()
	registros, err := conexionEstablecida.Query("SELECT * FROM actores")
	if err != nil {
		panic(err.Error())
	}
	actor := Actor{}
	arrActor := []Actor{}

	for registros.Next() {
		var id int
		var nombre string
		var edad int
		err = registros.Scan(&id, &nombre, &edad)
		if err != nil {
			panic(err.Error())
		}
		actor.id = id
		actor.nombre = nombre
		actor.edad = edad

		arrActor = append(arrActor, actor)

	}
	//fmt.Println(arrActor)

	plantillas.ExecuteTemplate(w, "inicio", arrActor)

}

func Actores(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hola mundo.")
	plantillas.ExecuteTemplate(w, "crear", nil)
}

func Insertar(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nombre := r.FormValue("nombre")
		edad := r.FormValue("edad")

		conexionEstablecida := conexionBD()

		insertarRegistros, err := conexionEstablecida.Prepare("INSERT INTO actores(nombre,edad) VALUES (?,?)")

		if err != nil {
			panic(err.Error())
		}
		insertarRegistros.Exec(nombre, edad)

		http.Redirect(w, r, "/", 301)
	}

}
