package handlers

import (
	"log"
	"net/http"
	"time"
)

func CommonHeadersHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/vnd.api+json")

		next.ServeHTTP(w, r)
	})
}

func LoggingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		format := "%s - - [%s] \"%s %s %s\" %s\n"
		log.Printf(format, r.RemoteAddr, time.Now().Format(time.RFC1123),
			r.Method, r.URL.Path, r.Proto, r.UserAgent())

		next.ServeHTTP(w, r)
	})
}
