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

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network addres")

	// Define a new command-line flag for the MySQL DSN string.
	dsn := flag.String("dsn", os.Getenv("SNIPPETBOX_DB_DSN"), "MySQL data source name")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// To keep the main() function tidy I've put the code for creating a connection
	// pool into the separate openDB() function below. We pass openDB() the DSN
	// from the command-line flag.
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	// We also defer a call to db.Close(), so that the connection pool is closed
	// before the main() function exits.
	defer db.Close()

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		// Call the new app.routes() method to get the servemux containing our routes.
		Handler: app.routes(),
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

// we create this function to save in the rootCertPool the CA vertificate
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
