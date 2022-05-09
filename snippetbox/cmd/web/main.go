package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
  
  _"github.com/go-sql-driver/mysql"
  "github.com/Suellen-Kitten/PF_CC5M-WEB/pkg/models/mysql"
)

type application struct{
  errorLog *log.Logger
  infoLog *log.Logger
  snippets *mysql.SnippetModel
}

//curl -i -X GET http://localhost:4000/snippetbox/create
func main() {
  //nome da flag, valor padrão e drescrição
  addr := flag.String("addr", ":4000", "Porta da Rede")
  dsn := flag.String("dsn",  
                     "qx8M4jivY9:X0shBFMXzG@tcp(remotemysql.com)/qx8M4jivY9?parseTime=true", 
                     "MySql DSN")
  
  flag.Parse()

  infoLog := log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)
  errorLog := log.New(os.Stderr, "ERRO:\t", log.Ldate|log.Ltime|log.Lshortfile)

  db, err := openDB(*dsn)
  if err != nil {
    errorLog.Fatal(err)
  }
  
  defer db.Close()

  app := &application{
    errorLog: errorLog,
    infoLog: infoLog,
    snippets: &mysql.SnippetModel{DB:db},
  }

  /* no Visual Code (ctrl + .) cria func
	mux := http.NewServeMux()
  
  mux.HandleFunc("/", app.home)
  mux.HandleFunc("/snippetbox", app.showSnippet)
  mux.HandleFunc("/snippetbox/create", app.createSnippet)

  fileServer := http.FileServer(http.Dir("./ui/static/"))
  mux.Handle("/static/",http.StripPrefix("/static",fileServer))*/

  srv := &http.Server{
    Addr: *addr,
    ErrorLog: errorLog,
    Handler: app.routes(),
  }
  
  infoLog.Printf("Inicializando o servidor na porta: %s\n", *addr)
  err = srv.ListenAndServe()
  errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error){
  db, err := sql.Open("mysql", dsn)
  if err != nil{
    return nil, err
  }
  if err = db.Ping(); err != nil{
    return nil, err 
  }
  return db, nil
}
