package autoRecover

import (
	"log"
	"net/http"
	"runtime/debug"
)

func SafeWrapper(req *http.Request, w http.ResponseWriter, f func(req *http.Request, w http.ResponseWriter)) {
	defer func() {
		p := recover()
		if p != nil {
			log.Printf("WARN: panic fired in %v.panic - %v\n", f, p)
			log.Println(string(debug.Stack()))
			w.WriteHeader(597)
		}
	}()
	f(req, w)
}
