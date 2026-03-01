package main

import (
	"log"
	"net/http"
	"time"

	"github.com/sunatsirt/glowing-chainsaw/internal/app"
	"github.com/sunatsirt/glowing-chainsaw/internal/config"
	"github.com/sunatsirt/glowing-chainsaw/internal/db"
)

func main() {
	cfg := config.Load()
	pool := db.NewPool(cfg.DatabaseURL)
	defer pool.Close()

	srv := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           app.NewRouter(pool),
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("listening on http://localhost:%s", cfg.Port)
	log.Fatal(srv.ListenAndServe())
}
