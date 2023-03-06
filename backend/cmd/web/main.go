package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	db *mongo.Client
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "mongodb://localhost:27017/?timeoutMS=10000", "Mongodb data source name")

	flag.Parse()

	// Use log.New() to create a logger for writing information messages. This takes
	// three parameters: the destination to write the logs to (os.Stdout), a string
	// prefix for message (INFO followed by a tab), and flags to indicate what
	// additional information to include (local date and time). Note that the flags
	// are joined using the bitwise OR operator |.
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	infoLog.Println("Connecting to MongoDB...[it can take up to 10 seconds]")
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	infoLog.Println("Pinged your deployment. You successfully connected to MongoDB!")

	// Disconnect the client once the function returns
	defer func() {
		if err = db.Disconnect(context.TODO()); err != nil {
			errorLog.Fatal(err)
		}
	}()

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		db: db,
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

func openDB(dsn string) (*mongo.Client, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(dsn).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		// prefer to lift up the error to the caller rather than handling it here
		return nil, err
	}

	// Send a ping to confirm a successful connection
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return nil, err
	}

	return client, nil
}
