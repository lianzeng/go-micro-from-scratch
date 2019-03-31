1. run "go get -u github.com/globalsign/mgo" to download mgo driver to $GOPATH/src ,  $GOPATH="~/go"

2. compile: go build simple_server.go mgoUtil.go 

3. start mongodb:  mongod --port 27017 --dbpath ~/demo/db

4.run demo:  ./simple_server -f config.conf 

5.check db result:  
5.1. login by mongo shell: mongo

5.2. use db_demo ; db.coll_demo.find(); 