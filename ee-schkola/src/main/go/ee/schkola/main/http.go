package main

import (
	"github.com/eugeis/gee/lg"
	"net/http"
	"github.com/gorilla/mux"
	"ee/schkola/person"
	"fmt"
	"github.com/eugeis/gee/net"
)

var log = lg.NewLogger("Schkola ")

func main() {
	log.Info("Server started")

	router := mux.NewRouter().StrictSlash(true)

	router.Methods(net.GET).Path("/").Name("Index").HandlerFunc(Index)

	personRouter := person.NewPersonRouter("")
	personRouter.Setup(router)

	log.Err("%v", http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
