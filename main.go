package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/oustaa/go-url-shortner/internal/routes"
)

func main() {
	port := int16(*flag.Int("port", 9000, "this is the port where the app will be runing"))
	flag.Parse()

	r := routes.GetRoutes()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}
