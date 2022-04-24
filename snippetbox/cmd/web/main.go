package main

import (
	"log"
  "flag"
  "net/http"
  "os"
)

type application struct{
  errorLog *log.Logger
  infoLog *log.Logger
}

//curl -i -X GET http://localhost:4000/snippetbox/create
func main() {
  //nome da flag, valor padrão e drescrição
  addr := flag.String("addr", ":4000", "Porta da Rede")
  flag.Parse()

  infoLog := log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)
  errorLog := log.New(os.Stderr, "ERRO:\t", log.Ldate|log.Ltime|log.Lshortfile)

  app := &application{
    errorLog: errorLog,
    infoLog: infoLog,
  }

	mux := http.NewServeMux()
  
  mux.HandleFunc("/", app.home)
  mux.HandleFunc("/snippetbox", app.showSnippet)
  mux.HandleFunc("/snippetbox/create", app.createSnippet)

  fileServer := http.FileServer(http.Dir("./ui/static/"))
  mux.Handle("/static/",http.StripPrefix("/static",fileServer))
  
  infoLog.Printf("Inicializando o servidor na porta: %s\n", *addr)
  err := http.ListenAndServe(*addr, mux)
  errorLog.Fatal(err)
}