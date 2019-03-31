package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type ErrorResp struct {
	Error string `json:"error"`
}

func Fatal(v ...interface{}) {
	os.Stderr.WriteString(fmt.Sprint(v, "\n"))
	os.Exit(1)
}

func ReplyWith(w http.ResponseWriter, code int, data interface{}) {
	var msg []byte
	if data != nil {
		msg, _ = json.Marshal(data)
	}
	h := w.Header()
	h.Set("Content-Length", strconv.Itoa(len(msg)))
	h.Set("Content-Type", "application/json")
	h.Set("Cache-Control", "no-store")
	w.WriteHeader(code)
	w.Write(msg)
}

func ReplyErr(w http.ResponseWriter, code int, err error) {
	ReplyWith(w, code, ErrorResp{err.Error()})
}
