package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type yell struct {
	What  string `json:"what"`
	Cloud string `json:"cloud"`
}

type response struct {
	Msg   string `json:"msg"`
	Cloud string `json:"cloud,omitempty"`
}

// Handler handles all incoming HTTP requests to this endpoint
func Handler(w http.ResponseWriter, r *http.Request) {
	// Ensure the request is a POST
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, getResponse(response{
			Msg: fmt.Sprintf("Method %s not supported", r.Method),
		}))
		return
	}

	y, err := getYellFromRequestBody(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, getResponse(response{
			Msg: "unable to parse request body",
		}))
	}

	if len(y.What) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, getResponse(response{
			Msg: "message not present",
		}))
	}

	if len(y.Cloud) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, getResponse(response{
			Msg: "select a cloud",
		}))
	}

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, getResponse(response{
		Msg:   y.What,
		Cloud: y.Cloud,
	}))
	return
}

func getYellFromRequestBody(body io.ReadCloser) (yell, error) {
	var y yell

	err := json.NewDecoder(body).Decode(&y)

	return y, err
}

func getResponse(concrete interface{}) string {
	b, _ := json.Marshal(concrete)

	return string(b)
}
