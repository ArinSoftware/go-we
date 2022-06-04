package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arinsoftware/go-web/pkg/config"
	"github.com/arinsoftware/go-web/pkg/handlers"
	"github.com/arinsoftware/go-web/pkg/render"
)

const portNum = ":4000"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Can not create template cache")
	}

	app.TemplateCache = tc

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s", portNum)
	_ = http.ListenAndServe(portNum, nil)
}
