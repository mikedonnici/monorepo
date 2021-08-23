package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type server struct {
	port uint
}

func main() {

	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {

	srvr := server{
		port: 8088,
	}

	http.HandleFunc("/", indexHandler())
	http.HandleFunc("/404", notFoundHandler())

	return http.ListenAndServe(fmt.Sprintf(":%d", srvr.port), nil)
}

func indexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, err := io.WriteString(w, "hello"); err != nil {
			log.Fatalf("could not write to ResponseWrite, err = %s", err)
		}
	}
}

func notFoundHandler() http.HandlerFunc {
	return http.NotFoundHandler().(http.HandlerFunc)
}
