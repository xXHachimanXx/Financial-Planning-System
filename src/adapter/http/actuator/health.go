package actuator

import (
	"encoding/json"
	"net/http"
)

type HealthBody struct {
	Status string `json:"status"`
}

func Health(rw http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	rw.Header().Set("Content-Type", "application/json")

	var status = HealthBody{"alive"}

	_ = json.NewEncoder(rw).Encode(status)
}
