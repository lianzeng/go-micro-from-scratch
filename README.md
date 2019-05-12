* if use vscode, need set go.gopath(within settings.json) to  "/Users/liangzeng/go/:path/to/current/dir"

*  run "go get -u github.com/globalsign/mgo" to download mongodb golang driver to $GOPATH/src ,  $GOPATH="~/go"
*  run "go get -u github.com/gorilla/mux" to download http route framework, it's bettern than native http.ServeMux.

* start mongodb:  mongod --port 27017 --dbpath ~/path/to/db

* compile: ./build.sh

* run server:  ./main -f config.conf 

* TODO: add test case(unit test + component test + benchmark test) , httptest.NewServer

* TODO: support  profile about cpu,memory for tracking perfomance problem

* TODO: support redis-cluster, redis-standalone

* TODO: support auditlog for recording req&resp.

* TODO: support nginx, nginx route api to go-server

* TODO: support monitor, use prometheus and grafana for metric-monitor

* TODO: support docker and kubernetes 

* TODO: support kafka(https://github.com/shopify/sarama)

* TODO: support Rate Limit