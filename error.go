package main

import (
	"fmt"
	"net/http"
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
