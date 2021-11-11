package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
)

func MovieShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie_id := params["id"]
	fmt.Fprintf(w, "Has cargado la pelicula numero %s",movie_id)
}