package AMJ

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	webart "AMJ/asciiart/MakeAsciiArt"
)

// FormHandler handles HTTP requests for the form submission.
func FormHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method == http.MethodPost {
		// Retrieve the value of the "textbox" form field
		textbox := r.FormValue("textbox")
		if textbox == "" {
			BadRequestErrorHandler(w, r)
			return
		}
	}

	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	data := PageData{false, "", "", "", "", ""}

	// Retrieve form values for textbox, color, style, and alignment
	str := r.FormValue("textbox")
	color := r.FormValue("color")
	font := r.FormValue("style")
	align := strings.TrimPrefix(r.FormValue("alignment"), "text-")

	// Generate ASCII art if the textbox is not empty
	if str != "" {
		output, err := webart.Asciiart(str, font)
		if err != nil {
			if strings.Contains(err.Error(), "400") {
				BadRequestErrorHandler(w, r)
				return
			} else if strings.Contains(err.Error(), "500") {
				InternalServerErrorHandler(w, r)
				return
			}
		}

		data = PageData{true, output, str, color, font, align}
	}

	// Parse and execute the "views/index.html" template
	tmpl := template.Must(template.ParseFiles("views/index.html"))
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
