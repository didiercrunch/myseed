package main

import (
	"log"
	"net/http"
	"fmt"
	"strconv"

	"github.com/gorilla/mux"
)

func createMuxRouter() http.Handler {
	r := mux.NewRouter()
	serveStatic(r.PathPrefix("/").Subrouter())
	return r
}

func main() {
	r := createMuxRouter()
	http.Handle("/", r)
	address := params.Host + ":" + strconv.Itoa(params.Port)
	fmt.Println("server running at", address)
	if err := http.ListenAndServe(address, createMuxRouter()); err != nil {
		log.Fatal(err)
	}
}
