package server

import (
	"fmt"
	"github.com/tiyee/palmon/engine"
	"github.com/tiyee/palmon/handle"
	"net/http"
)

func taskHandle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handle.Get(engine.NewContext(w, r))
	case "POST":
		handle.Post(engine.NewContext(w, r))
	}
}
func pingHandle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("pong !"))
}
func syncHandle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handle.Pull(engine.NewContext(w, r))
	case "POST":
		handle.Push(engine.NewContext(w, r))
	}
}
func Run() {
	http.HandleFunc("/task", taskHandle)
	http.HandleFunc("/sync", syncHandle)
	http.HandleFunc("/ping", pingHandle)
	if err := http.ListenAndServe(":8083", nil); err != nil {
		fmt.Println(err.Error())
	}
}
