package main

import (
	"net/http"
	"sync"
)

var bookstore map[string]string
var bookMutex sync.RWMutex

func main() {
	bookstore = make(map[string]string)
	bookMutex = sync.RWMutex{}
	http.HandleFunc("/get", get)
	http.HandleFunc("/set", set)
	http.HandleFunc("/remove", remove)
	http.HandleFunc("/list", list)
	http.ListenAndServe(":3000", nil)
}

func get(w http.ResponseWriter, r *http.Request) {
}

func set(w http.ResponseWriter, r *http.Request) {
}

func remove(w http.ResponseWriter, r *http.Request) {
}

func list(w http.ResponseWriter, r *http.Request) {
}
