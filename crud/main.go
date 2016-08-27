package crud

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
