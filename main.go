package main

import (
	"floodinskiy/handlers"
	"fmt"
	"net/http"
)

var port = ":9000"

func main() {
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/send", handlers.SendHandler)

	fmt.Printf("Server started at http://localhost:%s\n", port)
	http.ListenAndServe(":"+port, nil)
}
