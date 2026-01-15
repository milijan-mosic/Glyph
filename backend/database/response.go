package database

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, status int, key string, payload any) {
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(map[string]any{
		key: payload,
	})
}
