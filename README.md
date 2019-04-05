* if use vscode, need set go.gopath(within settings.json) to  "/Users/liangzeng/go/:path/to/cur/dir"

*  run "go get -u github.com/globalsign/mgo" to download mongodb golang driver to $GOPATH/src ,  $GOPATH="~/go"
*  run "go get -u github.com/gorilla/mux" to download http route framework, it's bettern then native http.ServeMux.

* start mongodb:  mongod --port 27017 --dbpath ~/demo/db

* compile: ./build.sh

* run server:  ./main -f config.conf 

* use curl as client:
* curl -I  -X POST http://127.0.0.1:8000/zhangshan/engineer
* curl   -X GET http://127.0.0.1:8000/zhangshan

* TODO: add test case , httptest.NewServer

* TODO: add  profile about cpu,memory for tracking perfomance problem

* TODO: add redis client , support redis-cluster, redis-standalone

* TODO: auto generate auditlog for recording req&resp.

* TODO: support nginx, nginx route api to go-server

* TODO: support prometheus and grafana for metric-monitor

* TODO: support docker and kubernetes 