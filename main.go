package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/idiomatic/locale"
	"github.com/idiomatic/middleware"
)

const (
	localizedDir = "localized"
	staticDir    = "static"
)

var locales = []string{"en-US", "en-CA", "en-GB", "fr-FR", "fr-FR"}

func main() {
	matcher := locale.NewMatcher(locales)

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong\n")
	})

	h := http.Handler(http.DefaultServeMux)
	h = middleware.FileServer(http.Dir(staticDir), h)
	h = locale.StripLocalePrefix(h)
	h = middleware.FileServer(http.Dir(localizedDir), h)
	h = locale.Handler(locales, matcher, h)
	addr := ":" + os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(addr, h))
}
