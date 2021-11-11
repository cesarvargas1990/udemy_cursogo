package main

import (
	"net/http"
	"log"
	"fmt"
)

func MovieList(w http.ResponseWriter, r *http.Request) {
	var results []Movie
	err := collection.Find(nil).Sort("-_id").All(&results)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Resultados", results)
	}
	responseMovies(w,200,results)
}