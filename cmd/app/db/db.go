package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

//Connect
func Connect() *pgxpool.Pool {
	login, exists := os.LookupEnv("DB_USER")
	if !exists {
		log.Fatalf("No found DB_USER")
	}
	pass, exists := os.LookupEnv("DB_PASS")
	if !exists {
		log.Fatalf("No found DB_PASS")
	}
	url, exists := os.LookupEnv("DB_URL")
	if !exists {
		log.Fatalf("No found DB_URL")
	}
	dbName, exists := os.LookupEnv("DB_NAME")
	if !exists {
		log.Fatalf("No found DB_NAME")
	}
	conn, error := pgxpool.Connect(context.Background(), "postgres://"+login+":"+pass+"@"+url+":5432/"+dbName)
	if error != nil {
		log.Fatalf("Can't connect to db : %v", error)
	}
	return conn
}

//Close connect
func Close(conn *pgxpool.Pool) {
	conn.Close()
}
