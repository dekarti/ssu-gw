package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/dekarti/ssu-gw/handlers"
	"github.com/justinas/alice"
)

func main() {
	var httpAddr = flag.String("http", "0.0.0.0:8080", "HTTP service address")
	flag.Parse()
	log.Println("Starting server...")
	log.Printf("HTTP service listening on %s", *httpAddr)

	mux := http.NewServeMux()

	commonHandlers := alice.New(handlers.CommonHeadersHandler, handlers.LoggingHandler)

	mux.Handle("/hello", commonHandlers.ThenFunc(http.HandlerFunc(handlers.HelloHandler)))

	log.Fatal(http.ListenAndServe(*httpAddr, mux))
}
