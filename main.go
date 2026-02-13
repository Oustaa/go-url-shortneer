package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
	"github.com/oustaa/go-url-shortner/internal/config"
	"github.com/oustaa/go-url-shortner/internal/db"
	"github.com/oustaa/go-url-shortner/internal/models"
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

	db.AutoMigrate(&models.User{}, &models.URL{}, &models.URLStats{})

	r := routes.GetRoutes(db)

	catchAllClientRoutesHandler := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/index.html")
	}

	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "public"))
	r.Handle("/public/*", http.StripPrefix("/public/", http.FileServer(filesDir)))

	r.Get("/*", catchAllClientRoutesHandler)

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("Application is runing on port %d", port)
	log.Fatal(s.ListenAndServe())
}
