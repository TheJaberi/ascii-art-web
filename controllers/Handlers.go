package webart

import (
    "fmt"
    "html/template"
    "net/http"
    "strings"
    asciiart "webart/asciiart/MakeAsciiArt"
)

type PageData struct {
    ShowOutput bool
    Output     string
    Str string
}

func FormHandler(w http.ResponseWriter, r *http.Request) {

    if r.URL.Path != "/" {
        http.Error(w, "Error 404 \npage not found.", http.StatusNotFound)
        return
    }
    
    if r.Method == http.MethodPost {
    textbox := r.FormValue("textbox")

    if textbox == "" {
        http.Error(w, "Error 400 \nTextbox cannot be empty", http.StatusBadRequest)
        return
    }
    }

    err := r.ParseForm()
    if err != nil {
        fmt.Fprintf(w, "ParseForm() err: %v", err)
        return
    }
    data := PageData{false, "",""}
    str1 := r.FormValue("textbox")
 
    // replacing \r\n with \\n so our ascii art program can deal with it
    strarr := strings.Split(str1,"\r\n")
    str := strings.Join(strarr,"\\n")
    if str != "" {
        for _, r := range str {
            if r < 32 || r > 126 {
                http.Error(w, "Error 500 \nonly use english", http.StatusInternalServerError)
                return
            }
        }
        style := "asciiart/Fonts/"+r.FormValue("style")+".txt"
        option := []string{"", "", ""}
        output := asciiart.ReturnArt(str, "", option,style, 0) ///// fix the width
        if output == "Error" {
            http.Error(w, "Error 500 ", http.StatusInternalServerError)
            return
        }
        data = PageData{true, output,str1}
    }
    // fmt.Println(output)
    tmpl := template.Must(template.ParseFiles("views/index.html"))
    err = tmpl.Execute(w, data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
