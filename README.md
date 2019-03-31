1. run "go get -u github.com/globalsign/mgo" to download mgo driver to $GOPATH/src ,  $GOPATH="~/go"

2. compile: go build simple_server.go mgoUtil.go 

3. start mongodb:  mongod --port 27017 --dbpath ~/demo/db

4.run demo:  ./simple_server -f config.conf 

5.use curl as client:
curl -I  -X POST http://127.0.0.1:8000/zhangshan/engineer
curl   -X GET http://127.0.0.1:8000/zhangshan

