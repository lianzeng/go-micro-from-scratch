package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"

	"github.com/globalsign/mgo"
)

type DocItem struct {
	Tbl string `bson:"tbl,omitempty"`
	Uid uint32 `bson:"uid,omitempty"`
}

type Config struct {
	MaxProcs int           `json:"max_procs"`
	BindHost string        `json:"bind_host"`
	DbConfig MongoDbConfig `json:"mongo"`
}

var (
	NEWLINE    = []byte{'\n'}
	collection *mgo.Collection
)

func Fatal(v ...interface{}) {
	os.Stderr.WriteString(fmt.Sprint(v, "\n"))
	os.Exit(1)
}

func defaultHandle(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello Go ! \nGoodbye C++.")
}

func RegisterHandler(svrmux *http.ServeMux) {
	svrmux.HandleFunc("/", defaultHandle)
}

func main() {
	fmt.Println("simple server with mongodb")

	var config Config
	if err := ReadConfig(&config); err != nil {
		Fatal("ReadConfig Failed. ", err)
	}

	runtime.GOMAXPROCS(config.MaxProcs)

	dbSesstion := ConnectMongo(&config.DbConfig)
	defer dbSesstion.Close()

	fmt.Println(config.DbConfig.Database)

	collection = dbSesstion.DB(config.DbConfig.Database).C(config.DbConfig.CollectionName)

	svrmux := http.NewServeMux()
	RegisterHandler(svrmux)

	Fatal(http.ListenAndServe(config.BindHost, svrmux))
}
