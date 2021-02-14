package usecase

import (
	"encoding/json"
	"net/http"
)

// HelloRinda is Call Rinda
func HelloRinda(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Hello Rinda")
}
