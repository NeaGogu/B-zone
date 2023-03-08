package main

import (
	"bzone/backend/internal/models"
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
	zipCodeDbModel *models.ZipCodeDBModel
}

func main() {
	// addr := flag.String("addr", ":4000", "HTTP network address")
	// dsn := flag.String("dsn", "mongodb://test_network:27017/?timeoutMS=10000", "Mongodb data source name")
	addr, ok := os.LookupEnv("ADDR")
	if !ok {
		log.Println("ADDR environment variable is not set. Using default value: :4000")
		addr = ":4000"
	}

	// if this env variable is not set then the database cannot be used
	// so we should exit the program
	dsn, ok := os.LookupEnv("MONGO_URL")
	if !ok {
		log.Fatal("MONGO_URL environment variable is not set")
	}

	flag.Parse()

	// Use log.New() to create a logger for writing information messages. This takes
	// three parameters: the destination to write the logs to (os.Stdout), a string
	// prefix for message (INFO followed by a tab), and flags to indicate what
	// additional information to include (local date and time). Note that the flags
	// are joined using the bitwise OR operator |.
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	infoLog.Println("Connecting to MongoDB...[it can take up to 10 seconds]")
	db, err := openDB(dsn)
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
		// TODO remove hardcoded value
		zipCodeDbModel: &models.ZipCodeDBModel{ DB: db.Database("zipcodes") },
	}

	srv := &http.Server{
		Addr:     addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", addr)
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
