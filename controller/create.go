package controller

import (
	"encoding/json"
	"net/http"
	"project/model"
	"project/views"
)

func create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			data := views.PostRequest{}

			json.NewDecoder(r.Body).Decode(&data)

			if err := model.Create(data.Fullname, data.Phone, data.City); err != nil {
				w.Write([]byte("Some error"))
				return
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(data)
		}
	}
}
