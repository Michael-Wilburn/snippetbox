package main

import (
	// New import

	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/Michael-Wilburn/snippetbox/internal/models"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

// Add a snippets field to the application struct. This will allow us to
// make the SnippetModel object available to our handlers.
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *models.SnippetModel
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network addres")
	dsn := flag.String("dsn", os.Getenv("SNIPPETBOX_DB_DSN"), "MySQL data source name")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	// Initialize a models.SnippetModel instance and add it to the application
	// dependencies.
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		snippets: &models.SnippetModel{DB: db},
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func (app *application) init() {
	rootCertPool := x509.NewCertPool()
	pem, err := os.ReadFile("database/cert/ca.pem")
	if err != nil {
		app.errorLog.Fatalf("Failed to read CA certificate file at 'database/cert/ca.pem': %v", err)
	}
	if !rootCertPool.AppendCertsFromPEM(pem) {
		app.errorLog.Fatal("Failed to add CA certificate to the certificate pool")
	}

	mysql.RegisterTLSConfig("custom", &tls.Config{
		RootCAs: rootCertPool,
	})
}
