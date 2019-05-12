package auditlog

import (
	"autoRecover"
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	logBody []byte
	code    int
}

//make a copy to log
func (r *responseWriter) Write(content []byte) (n int, e error) {
	n, e = r.ResponseWriter.Write(content)
	copy(r.logBody, content)
	return
}

func (r *responseWriter) WriteHeader(code int) {
	r.ResponseWriter.WriteHeader(code)
	r.code = code
}

//log request and response to logfile
func HandleWithAuditlog(route http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		handlerWithLogRecord(req, w, route.ServeHTTP)
	})
}

func handlerWithLogRecord(req *http.Request, w http.ResponseWriter, f func(w http.ResponseWriter, req *http.Request)) {
	startTime := time.Now().UnixNano() / 100
	w1 := &responseWriter{
		ResponseWriter: w,
		logBody:        make([]byte, 100),
		code:           200,
	}

	autoRecover.Handle(req, w1, f) //call f(w1,req)

	endTime := time.Now().UnixNano() / 100
	b := bytes.NewBuffer(nil)
	b.WriteString("REQ\t")
	b.WriteString(strconv.FormatInt(startTime, 10))
	b.WriteByte('\t')
	b.WriteString(req.Method)
	b.WriteByte('\t')
	b.WriteString(req.URL.Path)
	b.WriteByte('\t')
	header, _ := json.Marshal(req.Header)
	b.Write(header)
	b.WriteByte('\t')
	code := strconv.Itoa(w1.code)
	b.WriteString(code)
	b.WriteByte('\t')
	b.Write(w1.logBody)
	b.WriteByte('\t')
	b.WriteString(strconv.FormatInt(endTime-startTime, 10))
	logToFile(b.Bytes())

}
