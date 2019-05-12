package autoRecover

import (
	"log"
	"net/http"
	"runtime/debug"
)

func Handle(req *http.Request, w http.ResponseWriter, f func(w http.ResponseWriter, req *http.Request)) {
	defer func() {
		p := recover()
		if p != nil {
			log.Printf("WARN: panic fired in %v.panic - %v\n", f, p)
			log.Println(string(debug.Stack()))
			w.WriteHeader(597)
		}
	}()
	f(w, req)
}
