package handler

import (
	"fmt"
	"net/http"
)

// Handler handles all incoming HTTP requests to this endpoint
func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	query := r.URL.Query()

	message, present := query["msg"]
	if !present || len(message[0]) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"msg\": \"message not present fucker\"}")
		return
	}

	cloud, present := query["cloud"]
	if !present || len(cloud[0]) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"msg\": \"select a cloud asshole\"}")
		return
	}

	fmt.Fprintf(w, "{\"msg\": \"%s\", \"cloud\": \"%s\"}", message[0], cloud[0])
}
