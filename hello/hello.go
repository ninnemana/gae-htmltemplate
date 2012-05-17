package hello

import (
    "html/template"
    "net/http"
    "fmt"
)

func init() {
    http.HandleFunc("/", root)
    http.HandleFunc("/sign", sign)
}

type foo struct{
    Title string
}

func root(w http.ResponseWriter, r *http.Request) {
    t, err := template.New("guestbook").ParseFiles("guestbook.html")
    d := &foo{Title:"example.org"}
	err = t.ExecuteTemplate(w, "guestbook.html", d)
	if err != nil {
		fmt.Println(err)
	}
}

type result struct{
	Content string
}
func sign(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("signtemplate").ParseFiles("signtemplate.html")

    d := &result{Content:r.FormValue("content")}

	err = t.ExecuteTemplate(w, "signtemplate.html", d)
	if err != nil {
		fmt.Println(err)
	}	
}