package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/joho/godotenv"
	"github.com/oustaa/go-url-shortner/internal/config"
	"github.com/oustaa/go-url-shortner/internal/db"
	"github.com/oustaa/go-url-shortner/internal/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, assuming environment variables are set externally")
	}

	port := int16(*flag.Int("port", 9000, "this is the port where the app will be runing"))
	flag.Parse()

	dbConfig, err := config.GetDBConfigFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	db, err := db.Open(dbConfig)
	if err != nil {
		log.Fatalf("Error Connecting to the db: %V", err)
	}

	r := routes.GetRoutes(db)

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}
