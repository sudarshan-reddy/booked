package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func gateWay(w http.ResponseWriter, r *http.Request) (url.Values, bool) {
	values, err := url.ParseQuery(r.URL.RawQuery)
	if errReturn(err, "", w) {
		return nil, true
	}
	if len(values.Get("key")) == 0 {
		if errReturn(nil, "Wrong input key", w) {
			return nil, true
		}
	}
	return values, false
}

func get(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		values, hasErrors := gateWay(w, r)
		if hasErrors {
			return
		}
		fmt.Println("in")
		bookMutex.RLock()
		value := bookstore[string(values.Get("key"))]
		bookMutex.RUnlock()
		if value != "" {
			fmt.Fprintf(w, value)
		} else {
			errReturn(nil, "No such value", w)
		}

	} else {
		errReturn(nil, "Only GET accepted", w)
	}
}

func set(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

	}
}

func remove(w http.ResponseWriter, r *http.Request) {
}

func list(w http.ResponseWriter, r *http.Request) {
}
