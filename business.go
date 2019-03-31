package main

import (
	"io"
	"net/http"

	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
)

func defaultHandle(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello Go ! \n")
}

func addPeople(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	err := collection.Insert(People{vars["Name"], vars["Job"]})
	if err != nil {
		ReplyErr(w, 500, err)
	} else {
		ReplyWith(w, 200, nil)
	}
}

func findPeople(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	var result People
	err := collection.Find(bson.M{"name": vars["Name"]}).One(&result)
	if err != nil {
		ReplyErr(w, 500, err)
	} else {
		ReplyWith(w, 200, result)
	}
}
