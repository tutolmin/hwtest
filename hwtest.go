package main

import (
    "html/template"
    "log"
    "io/ioutil"
    "net/http"
)

type Page struct {
    Version string
}

func loadPage() *Page {
    body, _ := ioutil.ReadFile("version")
    return &Page{Version: string(body)}
}

func handler(w http.ResponseWriter, r *http.Request) {
    p := loadPage()
    t, _ := template.ParseFiles("hwtest.html")
    t.Execute(w, p)
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
