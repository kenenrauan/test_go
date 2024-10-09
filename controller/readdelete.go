package controller

import (
	"encoding/json"
	"net/http"
	"project/model"
)

func readdelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			name := r.URL.Query().Get("name")
			if name == "" {
				data, err := model.ReadAll()
				if err != nil {
					w.Write([]byte(err.Error()))
				}

				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(data)
			} else {
				data, err := model.ReadByName(name)

				if err != nil {
					w.Write([]byte(err.Error()))
				}

				w.WriteHeader(http.StatusOK)

				json.NewEncoder(w).Encode(data)
			}
		} else if r.Method == http.MethodDelete {
			name := r.URL.Query().Get("name")

			err := model.DeleteEmployee(name)

			if err != nil {
				w.Write([]byte("Some Error"))
			}

			w.WriteHeader(http.StatusOK)

			json.NewEncoder(w).Encode(struct {
				Status string `json:status`
			}{"Item deleted"})

		}
	}
}
