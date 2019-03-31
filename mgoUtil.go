package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/globalsign/mgo"
)

// run "go get -u github.com/globalsign/mgo" to download mgo code to $GOPATH/src
type MongoDbConfig struct {
	Host           string `json:"host"`
	Database       string `json:"db"`
	CollectionName string `json:"coll"`
	Mode           string `json:"mode"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	SyncTimeoutSec int64  `json:"timeout"` // 以秒为单位
	Pool           int    `json:"pool"`
}


func ConnectMongo(dbConfig *MongoDbConfig) *mgo.Session {

	info := &mgo.DialInfo{
		Addrs:     strings.Split(dbConfig.Host, ","),
		Database:  dbConfig.Database,
		Username:  dbConfig.Username,
		Password:  dbConfig.Password,
		Timeout:   time.Duration(dbConfig.SyncTimeoutSec) * time.Second * 2,
		PoolLimit: dbConfig.Pool,
	}

	session, err := mgo.DialWithInfo(info)
	if err != nil {
		panic(err)
	}
	fmt.Println("connect to mongodb ok.")
	// set session mode
	switch dbConfig.Mode {
	case "Strong":
		session.SetMode(mgo.Strong, true)
	case "Monotonic":
		session.SetMode(mgo.Monotonic, true)
	case "Eventual":
		session.SetMode(mgo.Eventual, true)
	default:
		session.SetMode(mgo.Strong, true)
	}
	session.SetSyncTimeout(time.Duration(dbConfig.SyncTimeoutSec) * time.Second)
	return session
}
