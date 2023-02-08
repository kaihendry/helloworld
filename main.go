package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/apex/gateway/v2"
	"golang.org/x/exp/slog"
)

var (
	GoVersion = runtime.Version()
)

func main() {

	logger := slog.New(slog.NewJSONHandler(os.Stdout))
	slog.SetDefault(logger.With("version", os.Getenv("version"), "goversion", GoVersion))

	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	wrappedMux := NewLogger(mux)

	var err error
	if _, ok := os.LookupEnv("AWS_EXECUTION_ENV"); ok {
		err = gateway.ListenAndServe("", wrappedMux)
	} else {
		slog.Info("local development", "port", os.Getenv("PORT"))
		err = http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), wrappedMux)
	}
	slog.Error("error listening", err)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-Version", fmt.Sprintf("Version: %s GoVersion: %s", os.Getenv("version"), GoVersion))
	fmt.Fprintf(w, "https://github.com/kaihendry/helloworld %s\n", time.Now())
}
