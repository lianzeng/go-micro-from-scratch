package main

import (
	"auditlog"
	"fmt"
	"net/http"
	"runtime"

	. "service"

	"github.com/globalsign/mgo"
	"github.com/gorilla/mux"
)

type Config struct {
	MaxProcs int           `json:"max_procs"`
	BindHost string        `json:"bind_host"`
	DbConfig MongoDbConfig `json:"mongo"`
}

var (
	mgoColl     *mgo.Collection
	PepoleIndex = mgo.Index{
		Key:    []string{"name"},
		Unique: true,
	}
)

func main() {
	fmt.Println("Hello, go server !")

	var config Config
	if err := ReadConfig(&config); err != nil {
		Fatal("ReadConfig Failed. ", err)
	}

	runtime.GOMAXPROCS(config.MaxProcs) //set max thread

	dbSesstion := ConnectMongo(&config.DbConfig)
	defer dbSesstion.Close()

	mgoColl = dbSesstion.DB(config.DbConfig.Database).C(config.DbConfig.CollectionName)
	if err := mgoColl.EnsureIndex(PepoleIndex); err != nil {
		Fatal("create db index failed.", err)
	}

	route := mux.NewRouter()
	svr := NewService(mgoColl)
	svr.RegisterHandler(route) //register restful API

	fmt.Println("start server ", config.BindHost)
	Fatal(http.ListenAndServe(config.BindHost, auditlog.HandleWithAuditlog(route))) //shouldn't return.
	//TODO: add  profile for cpu, memory
}
