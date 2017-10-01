package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", hello)

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatalf("error listening: %s", err)
	}

}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Hello string `json:"msg"`
	}{Hello: "Hello Youtubers"})

}
