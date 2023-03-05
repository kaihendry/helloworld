package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/apex/gateway/v2"
	"github.com/kaihendry/slogresponse"
	log "golang.org/x/exp/slog"

	_ "net/http/pprof"
)

var GoVersion = runtime.Version()

//go:embed static
var static embed.FS

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	mux.HandleFunc("/debug/pprof/", http.DefaultServeMux.ServeHTTP)
	wrappedMux := slogresponse.New(mux)
	var err error
	if _, ok := os.LookupEnv("AWS_LAMBDA_FUNCTION_NAME"); ok {
		log.SetDefault(log.New(log.NewJSONHandler(os.Stdout)))
		err = gateway.ListenAndServe("", wrappedMux)
	} else {
		log.SetDefault(log.New(log.NewTextHandler(os.Stdout)))
		log.Info("local development", "port", os.Getenv("PORT"))
		err = http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), wrappedMux)
	}
	log.Error("error listening", err)
}

// slow function
func slow() {
	log.Warn("slow function")
	time.Sleep(1 * time.Second)
}

func hello(w http.ResponseWriter, r *http.Request) {
	// slow()
	w.Header().Set("X-Version", fmt.Sprintf("%s %s", os.Getenv("version"), GoVersion))
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t := template.Must(template.New("").ParseFS(static, "static/index.html"))
	t.ExecuteTemplate(w, "index.html", struct {
		Env []string
	}{
		Env: filterAWSsecrets(os.Environ()),
	})
}

func filterAWSsecrets(env []string) []string {
	var filtered []string
	for _, e := range env {
		if !strings.HasPrefix(e, "AWS_SE") {
			filtered = append(filtered, e)
		}
	}
	return filtered
}
