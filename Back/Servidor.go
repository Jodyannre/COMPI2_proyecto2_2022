package main

import (
	"Back/principal"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mensaje desde GET")
}

func PostUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mensaje desde POST")
}

func (res *Respuesta) GetEntrada(w http.ResponseWriter, r *http.Request) {
	var respuesta Respuesta
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error en algo")
	}
	json.Unmarshal(reqBody, &respuesta)
	//Aqui ya tengo la cadena de entrada
	fmt.Println(respuesta.Contenido)
	res.Contenido = respuesta.Contenido
}

func (res *Respuesta) getConsola(w http.ResponseWriter, r *http.Request) {
	res.Respuesta = principal.EjecutarPrograma(res.Contenido)
	fmt.Println("Entro al get consola")
	fmt.Fprintf(w, res.Respuesta)
}

type Respuesta struct {
	Contenido string
	Respuesta string
}

var respuesta Respuesta

func main() {
	/*
		StrictSlash es para que la slash del final no afecte a la ruta
	*/

	//Crear el servidor
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/users", GetUsers).Methods("GET")
	r.HandleFunc("/api/users", PostUsers).Methods("POST")
	r.HandleFunc("/parse", respuesta.GetEntrada).Methods("POST")
	r.HandleFunc("/getConsola", respuesta.getConsola).Methods("GET")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./Web/")))
	server := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println("Servidor escuchando.......")
	server.ListenAndServe()
}
