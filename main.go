package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"runtime/debug"

	"github.com/apex/gateway/v2"
)

func main() {
	http.HandleFunc("GET /", hello)
	var err error
	if _, ok := os.LookupEnv("AWS_LAMBDA_FUNCTION_NAME"); ok {
		slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
		err = gateway.ListenAndServe("", nil)
	} else {
		slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, nil)))
		slog.Info("local development", "port", os.Getenv("PORT"))
		err = http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	}
	slog.Error("error listening", "error", err)
}

func hello(w http.ResponseWriter, r *http.Request) {
	buildInfo, ok := debug.ReadBuildInfo()
	if !ok {
		slog.Error("debug.ReadBuildInfo() failed")
		return
	}
	w.Header().Set("X-Version", buildInfo.Main.Version)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	response := map[string]string{"message": "hello world"}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		slog.Error("failed to encode response", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
