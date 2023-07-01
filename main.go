package main

import (
	"fmt"
	"net/http"

	webart "AMJ/controllers"
)

func main() {
	http.HandleFunc("/css/", webart.CSSHandler)
	http.HandleFunc("/", webart.Handler)
	http.HandleFunc("/ascii-art", webart.FormHandler)
	fmt.Println("Server listening on port http://localhost:8080 ...")
	http.ListenAndServe(":8080", nil)
}
