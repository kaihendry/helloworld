package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {

	http.HandleFunc("/", slow)

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatalf("error listening: %s", err)
	}

}

func slow(w http.ResponseWriter, r *http.Request) {
	i, err := strconv.Atoi(r.URL.Path[1:])
	if err != nil {
		i = 0
	}

	after := time.Duration(i)

	log.Println("Sleep for", i)

	time.Sleep(after * time.Second)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Delay int `json:"delay"`
	}{Delay: i})

}
