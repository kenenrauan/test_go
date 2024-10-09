package controller

import (
	"net/http"
)

func Register() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/create", create())
	mux.HandleFunc("/", readdelete())
	return mux
}
