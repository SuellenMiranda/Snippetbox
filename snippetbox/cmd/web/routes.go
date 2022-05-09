package main

import "net/http"

func (app *application) routes() *http.ServeMux {
  mux := http.NewServeMux()
  
  mux.HandleFunc("/", app.home)
  mux.HandleFunc("/snippetbox", app.showSnippet)
  mux.HandleFunc("/snippetbox/create", app.createSnippet)

  fileServer := http.FileServer(http.Dir("./ui/static/"))
  mux.Handle("/static/",http.StripPrefix("/static",fileServer))

  return mux
}