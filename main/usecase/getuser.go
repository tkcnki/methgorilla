package usecase

import (
	"encoding/json"
	"net/http"
	"tkcnki/methgollira/main/model/entity"

	"github.com/gorilla/mux"
)

// GetSingleUser is return One User
func GetSingleUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	key := vars["id"]

	for _, item := range entity.Users {
		if item.Id == key {
			json.NewEncoder(w).Encode(item)
		}
	}
}
