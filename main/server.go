package main


import (
    "fmt"
    "log"
    "net/http"
    "asciiart"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
    if err := r.ParseForm(); err != nil {
        fmt.Fprintf(w, "ParseForm() err: %v", err)
        return
    }
    fmt.Fprintf(w, "POST request successful\n")
    text := r.FormValue("text")
    subStr := r.FormValue("substring")
    font := r.FormValue("style")
    color := r.FormValue("color")
    alignment := r.FormValue("alignment")
    arr := asciiart.ProcessAscii(text, subStr, font, color, alignment)
    for _,str := range arr {
        fmt.Fprintf(w, (str+"\n"))
    }

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/hello" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }

    if r.Method != "GET" {
        http.Error(w, "Method is not supported.", http.StatusNotFound)
        return
    }


    fmt.Fprintf(w, "Hello!")
}


func main() {
    fileServer := http.FileServer(http.Dir("../static"))
    http.Handle("/", fileServer)
    http.HandleFunc("/form", formHandler)
    http.HandleFunc("/hello", helloHandler)


    fmt.Printf("Starting server at port 8080\n")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}