package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"project/model"
	"project/views"
)

func create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			data := views.PostRequest{}

			json.NewDecoder(r.Body).Decode(&data)
			fmt.Println(data.Fullname)
			fmt.Println(data.Phone)
			fmt.Println(data.City)

			if err := model.Create(data.Fullname, data.Phone, data.City); err != nil {
				w.Write([]byte("Some error"))
				return
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(data)

		}
	}

}
