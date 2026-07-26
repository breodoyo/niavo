package user

import (
	"encoding/json"
	"net/http"

)

type User struct {
	ID    int `json:"id"`
	Name  string `json:"name"`
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	users := []User {
		{
			ID: 1,
			Name: "Brender",
		},
		{
			ID: 2,
			Name: "Mary",
		},
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)


	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "failed to encode users", http.StatusInternalServerError)
	}

}