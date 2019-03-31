*  run "go get -u github.com/globalsign/mgo" to download mgo driver to $GOPATH/src ,  $GOPATH="~/go"

* compile: go build simple_server.go mgoUtil.go 

* start mongodb:  mongod --port 27017 --dbpath ~/demo/db

* run demo:  ./simple_server -f config.conf 

* use curl as client:
* curl -I  -X POST http://127.0.0.1:8000/zhangshan/engineer
* curl   -X GET http://127.0.0.1:8000/zhangshan

