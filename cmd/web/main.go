package main

import (
	"fmt"
	"github.com/krishmagar/go-api/pkg/config"
	"github.com/krishmagar/go-api/pkg/handlers"
	"net/http"
)

const PORT string = ":8080"

func main() {
	var app config.AppConfig

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port%s", PORT))
	http.ListenAndServe(PORT, nil)
}
