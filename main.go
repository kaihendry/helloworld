package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"
)

func main() {

	http.HandleFunc("/fast", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintln(w, "fast") })

	http.HandleFunc("/slow", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Duration(2) * time.Second)
		fmt.Fprintln(w, "slow")
	})

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatalf("error listening: %s", err)
	}

}
