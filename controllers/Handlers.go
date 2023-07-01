package AMJ

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	webart "AMJ/asciiart/MakeAsciiArt"
)

type PageData struct {
	ShowOutput bool
	Output     string
	Str        string
	Color      string
	Font       string
	Alignment  string
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		textbox := r.FormValue("textbox")
		if textbox == "" {
			BadRequestErrorHandler(w, r)
			// http.Error(w, "Error 400: Textbox cannot be empty", http.StatusBadRequest)
			return
		}
	}

	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	data := PageData{false, "", "", "", "", ""}

	str := r.FormValue("textbox")
	color := r.FormValue("color")
	font := r.FormValue("style")
	align := strings.TrimPrefix(r.FormValue("alignment"), "text-")
	if str != "" {
		output, err := webart.Asciiart(str, font)
		// fmt.Println(output)
		if err != nil {
			if strings.Contains(err.Error(), "400") {
				BadRequestErrorHandler(w, r)
				// http.Error(w, "Error 400:\nBad Request: Use English chars only", http.StatusBadRequest)
				return
			} else if strings.Contains(err.Error(), "500") {
				InternalServerErrorHandler(w, r)
				// http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
				return
			}
		}

		data = PageData{true, output, str, color, font, align}
	}
	tmpl := template.Must(template.ParseFiles("views/index.html"))
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CSSHandler(w http.ResponseWriter, r *http.Request) {
	filePath := "views" + r.URL.Path

	http.ServeFile(w, r, filePath)
}

func BadRequestErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	if r.FormValue("textbox") != "" {
		http.ServeFile(w, r, "views/errorpages/eng400.html")
	} else {
		http.ServeFile(w, r, "views/errorpages/400.html")
	}
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	http.ServeFile(w, r, "views/errorpages/404.html")
}

func InternalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	http.ServeFile(w, r, "views/errorpages/500.html")
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" && r.URL.Path != "ascii-art" && r.URL.Path != "/css/" {
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
