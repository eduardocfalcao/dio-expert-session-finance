package actuator

import (
	"encoding/json"
	"net/http"
)

type Status struct {
	Health string `json:"health"`
}

func Healthcheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "application/json")

	var status = Status{
		Health: "alive",
	}
	
	_ = json.NewEncoder(w).Encode(status)
}
