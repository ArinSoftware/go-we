package main

import (
	"fmt"
	"net/http"

	"github.com/arinsoftware/go-web/pkg/handlers"
)

const portNum = ":4000"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s", portNum)
	_ = http.ListenAndServe(portNum, nil)
}
