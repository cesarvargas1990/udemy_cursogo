package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"	
)

func MovieUpdate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie_id := params["id"]
	

	if !bson.IsObjectIdHex(movie_id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(movie_id)
	decoder := json.NewDecoder(r.Body)
	
	var movie_data Movie
	err  := decoder.Decode(&movie_data)

	if err != nil {
		panic(err)
		w.WriteHeader(500)
		return
	}

	defer r.Body.Close()

	document := bson.M{"_id": oid}
	change := bson.M{"$set" : movie_data}
	err = collection.Update(document, change)

	if err != nil {
		panic(err)
		w.WriteHeader(500)
		return
	}
	
	responseMovie(w,200,movie_data)
}