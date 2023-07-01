package AMJ

import (
	"html/template"
	"net/http"
)

type PageData struct {
	ShowOutput bool
	Output     string
	Str        string
	Color      string
	Font       string
	Alignment  string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		NotFoundHandler(w, r)
		// http.Error(w, "Error 404: Page Not Found", http.StatusNotFound)
		return
	}

	if r.Method == "GET" {
		tmpl, err := template.ParseFiles("views/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
