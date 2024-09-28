package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/tokha04/swe-farmer-market-system/api"
	db "github.com/tokha04/swe-farmer-market-system/db/sqlc"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("could not load .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	conn, err := pgx.Connect(context.Background(), "postgresql://root:changeme@localhost:5432/fms?sslmode=disable")
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	defer conn.Close(context.Background())

	q := db.New(conn)

	router := gin.Default()
	api.UserRoutes(router, q)
	router.Use(api.Authorization())
	api.FarmerRoutes(router, q)

	log.Fatal(router.Run(":" + port))
}
