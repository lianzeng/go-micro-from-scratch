*  run "go get -u github.com/globalsign/mgo" to download mongodb golang driver to $GOPATH/src ,  $GOPATH="~/go"
*  run "go get -u github.com/gorilla/mux" to download http route framework, it's bettern then native http.ServeMux.

* compile: go build simple_server.go mgoUtil.go 

* start mongodb:  mongod --port 27017 --dbpath ~/demo/db

* run demo:  ./simple_server -f config.conf 

* use curl as client:
* curl -I  -X POST http://127.0.0.1:8000/zhangshan/engineer
* curl   -X GET http://127.0.0.1:8000/zhangshan

