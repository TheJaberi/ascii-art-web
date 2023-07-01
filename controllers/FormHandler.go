package AMJ

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	webart "AMJ/asciiart/MakeAsciiArt"
)

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
