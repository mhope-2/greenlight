package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

// Config struct
type Config struct {
	port int
	env string
	db struct {
		dsn string
	}
}

// Application struct
type Application struct {
	config Config
	logger *log.Logger
}


func main() {

	fmt.Println(os.Getenv("GREENLIGHT_DB_DSN"))

	// Declare an instance of the config struct.
	var cfg Config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")

	// Read the DSN value from the db-dsn command-line flag into the config struct. We
	// default to using our development DSN if no flag is provided.
	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("GREENLIGHT_DB_DSN"), "PostgreSQL DSN")

	flag.Parse()

	// Initialize a new logger which writes messages to the standard out stream,
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// connect to db
	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}

	// Defer a call to db.Close() so that the connection pool is closed before the
	//main() function exits.
	defer db.Close()

	// Also log a message to say that the connection pool has been successfully // established.
	logger.Printf("database connection pool established")


	// Declare an instance of the application struct, containing the config struct and the logger.
	app := &Application{
		config: cfg,
		logger: logger,
	}


	// Use the httprouter instance returned by app.routes() as the server handler.
	 srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}

	// Start the HTTP
	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err = srv.ListenAndServe()
	logger.Fatal(err)

}

// The openDB() function returns a sql.DB connection pool.
func openDB(cfg Config) (*sql.DB, error) {
	// Use sql.Open() to create an empty connection pool, using the DSN from the config
	//struct.
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	// Create a context with a 5-second timeout deadline.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Use PingContext() to establish a new connection to the database
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	// Return the sql.DB connection pool.
	return db, nil
}