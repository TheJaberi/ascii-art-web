package main

import (
	webart "webart/controllers"
	"fmt"
	"net/http"
)

func main() {
	// http.HandleFunc("/", webart.Handler)
	http.HandleFunc("/", webart.FormHandler)
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
