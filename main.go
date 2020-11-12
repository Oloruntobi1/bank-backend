package main

import (
	"database/sql"
	"log"
	"os"
	_"net/http"

	_"github.com/prometheus/client_golang/prometheus"
	_"github.com/prometheus/client_golang/prometheus/promhttp"
	_"github.com/lib/pq"
	"github.com/Oloruntobi1/bankBackend/api"
	db "github.com/Oloruntobi1/bankBackend/db/sqlc"
	"github.com/Oloruntobi1/bankBackend/rdstore"
	"github.com/go-redis/redis/v7"
	
	
)





func init() {
	//Initializing redis
	dsn := os.Getenv("REDIS_DSN")
	if len(dsn) == 0 {
		dsn = "localhost:6379"
	}
	rdstore.Client = redis.NewClient(&redis.Options{
		Addr: dsn, //redis port
	})
	_, err := rdstore.Client.Ping().Result()
	if err != nil {
		panic(err)
	}
}
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