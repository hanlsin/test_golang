package main

import (
	"fmt"
	"net/http"
)

type rootHandler struct {
}

func (h rootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello~!! This is a simple http server: %s\n", r.URL.Path)
}

func newHttpServer() *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "You hit foo!!")
	})

	mux.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "You hit bar!!")
	})

	mux.Handle("/", rootHandler{})

	/*
		listener, err := net.Listen("tcp", "0.0.0.0:4567")

		err = http.Serve(listener, mux)
		if err != nil {
			fmt.Println("ERROR: http.Serve:", err)
		}
	*/

	return &http.Server{Addr: ":4567", Handler: mux}
}

func StartHTTPServer(httpSrvr *http.Server) {
	err := httpSrvr.ListenAndServe()
	if err != nil {
		fmt.Println("ERROR: http.ListenAndServe:", err)
	}
}
