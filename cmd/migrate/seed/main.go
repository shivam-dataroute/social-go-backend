package main

import (
	"log"

	"github.com/SHIVAM-GOUR/social_go_app/internal/db"
	"github.com/SHIVAM-GOUR/social_go_app/internal/env"
	"github.com/SHIVAM-GOUR/social_go_app/internal/store"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/social?sslmode=disable")

	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	store := store.NewStorage(conn)
	db.Seed(&store)

}
