package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"	
)


func MovieDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie_id := params["id"]
	if !bson.IsObjectIdHex(movie_id) {
		w.WriteHeader(404)
		return
	}
	oid := bson.ObjectIdHex(movie_id)
	err := collection.RemoveId(oid)
	if err!= nil {
		w.WriteHeader(404)
		return
	} 
	results := Message{"success","La pelicula con ID "+movie_id+" ha sido eliminada correctamente"}
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(results)
}