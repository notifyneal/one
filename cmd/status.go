package cmd

import (
	"encoding/json"
	"net/http"
)

type HealthCheck struct {
	Status string `json:"status"`
}

func Status() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		status := HealthCheck{Status: "Ok"}
		json.NewEncoder(w).Encode(status)

	})
}
