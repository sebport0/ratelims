package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle(
		"GET /v1/tokenbucket/hello",
		tokenBucketMiddleware(http.HandlerFunc(v1TokenBucketHello)),
	)
	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)
}

func v1TokenBucketHello(w http.ResponseWriter, r *http.Request) {
	resp, _ := json.Marshal(map[string]string{"status": "success", "data": "hello token bucket!"})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func tooManyRequests(w http.ResponseWriter) {
	resp, _ := json.Marshal(map[string]string{"status": "error", "message": "too many requests"})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusTooManyRequests)
	w.Write(resp)
}
