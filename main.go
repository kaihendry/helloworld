package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/apex/gateway/v2"
	"golang.org/x/exp/slog"
)

var (
	Version   string
	GoVersion = runtime.Version()
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout)))
	log.Println("Version:", Version, "GoVersion:", GoVersion)

	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	wrappedMux := NewLogger(mux)

	if _, ok := os.LookupEnv("AWS_EXECUTION_ENV"); ok {
		log.Fatal(gateway.ListenAndServe("", wrappedMux), nil)
	} else {
		log.Printf("Assuming local development")
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), wrappedMux), nil)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-Version", fmt.Sprintf("Version: %s GoVersion: %s", Version, GoVersion))
	fmt.Fprintf(w, "https://github.com/kaihendry/helloworld %s\n", time.Now())
}
