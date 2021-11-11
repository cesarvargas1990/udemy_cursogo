package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

func MovieShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie_id := params["id"]
	if !bson.IsObjectIdHex(movie_id) {
		w.WriteHeader(404)
		return
	}
	oid := bson.ObjectIdHex(movie_id)
	results := Movie{}
	err := collection.FindId(oid).One(&results)
	if err!= nil {
		w.WriteHeader(404)
		return
	} 
	responseMovie(w,200,results)
}