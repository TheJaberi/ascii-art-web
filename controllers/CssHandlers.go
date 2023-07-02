package AMJ

import (
	"net/http"
)

func CSSHandler(w http.ResponseWriter, r *http.Request) {
	filePath := "views" + r.URL.Path
	http.ServeFile(w, r, filePath)
}
