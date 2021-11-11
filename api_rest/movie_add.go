package main

import (
	"net/http"
	"log"
	"encoding/json"
)

var movies = Movies{
	Movie{"destino fatal", 2013, "Sara konor"},
	Movie{"destino fatal 2", 2015, "Sara konor 2"},
	Movie{"movie fexample", 2021, "cesar konor 2"},
}

func MovieAdd(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var movie_data Movie 
	err := decoder.Decode(&movie_data)
	if (err !=nil) {
		panic(err)
	}
	defer r.Body.Close()
	log.Println(movie_data)
	movies = append(movies, movie_data)
}