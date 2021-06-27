package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

	// Save the yell off-thread b/c we don't care if it works
	go saveYell(y)

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, getResponse(response{
		Msg:   y.What,
		Cloud: y.Cloud,
	}))
	return
}

func saveYell(y yell) {
	logger := log.Default()

	apiKey := os.Getenv("SUPABASE_API_KEY")
	if len(apiKey) == 0 {
		logger.Print("no api key")
		return
	}

	endpoint := os.Getenv("SUPABASE_ENDPOINT")
	if len(endpoint) == 0 {
		logger.Print("no endpoint")
		return
	}

	body, _ := json.Marshal(y)

	request, err := http.NewRequest(
		http.MethodPost,
		endpoint,
		bytes.NewBuffer(body),
	)
	if err != nil {
		logger.Print(err)
		return
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Prefer", "return=representation")
	request.Header.Set("apikey", apiKey)

	client := &http.Client{}
	r, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		logger.Print(err)
		return
	}

	respBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		logger.Print(err)
		return
	}
	msg := fmt.Sprintf("supabase resp: %s", string(respBody))
	logger.Print(msg)
	fmt.Println(msg)
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
