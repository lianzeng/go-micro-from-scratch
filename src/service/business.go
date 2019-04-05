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
	route.HandleFunc("/{Name}/{Job}", s.addPeople).Methods("POST")
	route.HandleFunc("/{Name}", s.findPeople).Methods("GET")
	route.PathPrefix("/").HandlerFunc(defaultHandle) //default route
}

func defaultHandle(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello Go ! \n no handler register for path:"+req.URL.Path+"\n")
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
