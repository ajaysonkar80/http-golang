package main

import (
    "html/template"
    "net/http"
    "os"
)

type Todo struct {
    Title string
    Done  bool
}

type TodoPageData struct {
    PageTitle string
    Todos     []Todo
}

func main() {
    tmpl := template.Must(template.ParseFiles("layout.html"))
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        data := TodoPageData{
            PageTitle: "My TODO list",
            Todos: []Todo{
                {Title: "Task 1", Done: false},
                {Title: "Task 2", Done: true},
                {Title: "Task 3", Done: true},
            },
        }
        tmpl.Execute(w, data)
    })

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // Default to port 8080 if not set
    }

    http.ListenAndServe(":" + port, nil)
}
