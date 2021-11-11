package main

import (
	"net/http"
	"encoding/json"
)

func MovieList(w http.ResponseWriter, r *http.Request) {
	movies := Movies{
		Movie{"destino fatal", 2013, "Sara konor"},
		Movie{"destino fatal 2", 2015, "Sara konor 2"},
		Movie{"movie fexample", 2021, "cesar konor 2"},
	}
	json.NewEncoder(w).Encode(movies)
}