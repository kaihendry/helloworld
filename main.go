package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strconv"
	"time"
)

func main() {

	http.HandleFunc("/", root)

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatalf("error listening: %s", err)
	}

}

func root(w http.ResponseWriter, r *http.Request) {
	delay, err := strconv.Atoi(r.URL.Query().Get("delay"))
	if err == nil {
		time.Sleep(time.Duration(delay) * time.Second)
	}
	fmt.Fprintln(w, fmt.Sprintf("hallo, %d", delay))
}
