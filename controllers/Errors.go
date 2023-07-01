package AMJ

import (
	"net/http"
)

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
