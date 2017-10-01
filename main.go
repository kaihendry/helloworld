package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/json"
)

func main() {
	log.SetHandler(json.New(os.Stderr))

	http.HandleFunc("/", hello("Local"))
	http.HandleFunc("/news", hello("there"))

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatalf("error listening: %s", err)
	}

}

func hello(name string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		n := r.FormValue("name")
		if n != "" {
			name = n
		}
		fmt.Fprintln(w, "Hello "+name)

		// curl -X POST --data-urlencode "name=test" http://localhost:3000

	}
}
