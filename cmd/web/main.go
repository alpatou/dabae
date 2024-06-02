package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/alpatou/dabae/internal/models"
	_ "github.com/go-sql-driver/mysql" // New import
	"github.com/joho/godotenv"
)

type config struct {
	addr      string
	staticDir string
}

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *models.SnippetModel
}

func main() {

	envFileNotLoaded := godotenv.Load()

	if envFileNotLoaded != nil {
		log.Fatal("Error loading .env")
	}

	// addr := flag.String("addr", ":4000", "obvious")

	var cfg config

	flag.StringVar(&cfg.addr, "addr", ":4000", "Port")
	flag.StringVar(&cfg.staticDir, "static-dir", "./ui/static", "Path to static assets")

	dsn := flag.String("dsn", "web:pass@tcp(localhost:3306)/snippetbox?parseTime=true", "Mysql data source")

	flag.Parse()

	infolog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	errLog := log.New(os.Stderr, "Error\t", log.Ldate|log.Ltime|log.Lshortfile)
	db, db_err := openDB(*dsn)

	if db_err != nil {
		errLog.Fatal(db_err)
	}

	ping_err := db.Ping()

	if ping_err != nil {
		errLog.Fatal(ping_err)
	}
	// it causes access denied, well ping does to be honest
	defer db.Close()

	app := &application{
		errorLog: errLog,
		infoLog:  infolog,
		snippets: &models.SnippetModel{DB: db},
	}

	srv := &http.Server{
		Addr:     cfg.addr,
		ErrorLog: errLog,
		Handler:  app.routes(),
	}

	infolog.Printf("start serving on  %s", cfg.addr)
	err := srv.ListenAndServe()
	// log to a file someday
	errLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	// if err := db.Ping(); err != nil {
	// 	errLog.Fatal(err)
	// 	return nil, err
	// }
	return db, nil
}
