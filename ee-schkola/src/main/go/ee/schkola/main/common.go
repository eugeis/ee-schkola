package main

import (
	"github.com/eugeis/gee/lg"
	"net/http"
	"fmt"
)

var log = lg.NewLogger("Schkola ")

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
