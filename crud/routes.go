package crud

import (
	"fmt"
	"net/http"
	"net/url"
)

func errReturn(err error, errString string, w http.ResponseWriter) bool {
	if err != nil || errString != "" {
		w.WriteHeader(http.StatusBadRequest)
		if err != nil {
			fmt.Fprint(w, "Error: ", err)
		}
		fmt.Fprint(w, "Error: ", errString)
		return true
	}
	return false
}

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
		values, hasErrors := gateWay(w, r)
		if hasErrors {
			return
		}
		bookMutex.Lock()
		bookstore[string(values.Get("key"))] = string(values.Get("value"))
		bookMutex.Unlock()

		fmt.Fprintf(w, "success")
	} else {
		errReturn(nil, "Only POST accepted", w)
	}
}

func remove(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		values, hasErrors := gateWay(w, r)
		if hasErrors {
			return
		}
		bookMutex.Lock()
		delete(bookstore, values.Get("key"))
		bookMutex.Unlock()

		fmt.Fprint(w, "success")
	}
}

func list(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		bookMutex.RLock()
		for key, value := range bookstore {
			fmt.Fprintln(w, key, ":", value)
		}
		bookMutex.RUnlock()
	} else {
		errReturn(nil, "Only GET accepted", w)
	}
}
