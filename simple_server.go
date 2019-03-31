package main

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/gorilla/mux"

	"github.com/globalsign/mgo"
)

type People struct {
	Name string `bson:"name,omitempty" json:"name,omitempty"`
	Job  string `bson:"job,omitempty" json:"job,omitempty"`
}

type Config struct {
	MaxProcs int           `json:"max_procs"`
	BindHost string        `json:"bind_host"`
	DbConfig MongoDbConfig `json:"mongo"`
}

var (
	NEWLINE     = []byte{'\n'}
	collection  *mgo.Collection
	PepoleIndex = mgo.Index{
		Key:    []string{"name"},
		Unique: true,
	}
)

func RegisterHandler(route *mux.Router) {
	route.HandleFunc("/{Name}/{Job}", addPeople).Methods("POST")
	route.HandleFunc("/{Name}", findPeople).Methods("GET")
	route.PathPrefix("/").HandlerFunc(defaultHandle) //default route
}

func main() {
	fmt.Println("simple server with mongodb")

	var config Config
	if err := ReadConfig(&config); err != nil {
		Fatal("ReadConfig Failed. ", err)
	}

	runtime.GOMAXPROCS(config.MaxProcs) //set max thread

	dbSesstion := ConnectMongo(&config.DbConfig)
	defer dbSesstion.Close()

	collection = dbSesstion.DB(config.DbConfig.Database).C(config.DbConfig.CollectionName)
	if err := collection.EnsureIndex(PepoleIndex); err != nil {
		Fatal("create db index failed.", err)
	}

	route := mux.NewRouter()
	RegisterHandler(route)

	Fatal(http.ListenAndServe(config.BindHost, route))
}
