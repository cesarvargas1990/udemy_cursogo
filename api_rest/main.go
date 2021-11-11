package main

import ("net/http"
		"log"
	)

var collection = getSession().DB("movies").C("movies")

func main() {
	router := NewRouter()
	server := http.ListenAndServe(":8080", router)
	log.Fatal(server)
}