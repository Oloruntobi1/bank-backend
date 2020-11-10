package main

import (
	"database/sql"
	"log"
	"os"
	_"net/http"

	_"github.com/prometheus/client_golang/prometheus"
	_"github.com/prometheus/client_golang/prometheus/promhttp"
	_"github.com/lib/pq"
	"github.com/Oloruntobi1/Oloruntobi1/bank_backend/api"
	db "github.com/Oloruntobi1/Oloruntobi1/bank_backend/db/sqlc"
)

const serverAddress = "0.0.0.0:2000"
var (
	dbDriver = os.Getenv("DB_DRIVER")
	dbSource = os.Getenv("BB_SOURCE")
)

func main() {

	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("could not connect to the database", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	// http.Handle("/metrics", promhttp.Handler())

	//Here we are telling prometheus to keep track of ordersPlaced metric
	// prometheus.MustRegister(api.RequestsToCreateAccount)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}