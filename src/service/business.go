package service

import (
	"io"
	"net/http"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
)

type People struct {
	Name string `bson:"name,omitempty" json:"name,omitempty"`
	Job  string `bson:"job,omitempty" json:"job,omitempty"`
}

type Service struct {
	*mgo.Collection
}

func NewService(coll *mgo.Collection) *Service {
	return &Service{
		Collection: coll,
	}
}

func (s *Service) RegisterHandler(route *mux.Router) {
	route.HandleFunc("/panic", func(http.ResponseWriter, *http.Request) { panic("test panic auto recover.") })
	route.HandleFunc("/{Name}/{Job}", s.addPeople).Methods("POST")
	route.HandleFunc("/{Name}", s.findPeople).Methods("GET")
	route.PathPrefix("/").HandlerFunc(defaultHandle) //default route
}

func defaultHandle(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(501)
	io.WriteString(w, "not implemented for path "+req.URL.Path)
}

func (s *Service) addPeople(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	err := s.Insert(People{vars["Name"], vars["Job"]})
	if err != nil {
		ReplyErr(w, 500, err)
	} else {
		ReplyWith(w, 200, nil)
	}
}

func (s *Service) findPeople(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	var result People
	err := s.Find(bson.M{"name": vars["Name"]}).One(&result)
	if err != nil {
		ReplyErr(w, 500, err)
	} else {
		ReplyWith(w, 200, result)
	}
}
