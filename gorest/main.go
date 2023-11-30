package main

import (
	"context"
	"go-rest/db"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	// Initialize the Gin router
	r := gin.Default()
	ctx := context.Background()

	// Connect to the database

	dbpool, err := pgxpool.New(ctx, "postgres://go_user:secretpassword@localhost:5432/users_db")

	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer dbpool.Close()

	// Initialize the Queries struct from sqlc
	queries := db.New(dbpool)

	r.POST("/api/users", func(c *gin.Context) { CreateUserHandler(c, queries) })
	r.POST("/api/users/generateotp", func(c *gin.Context) { GenerateOTPHandler(c, queries) })
	r.POST("/api/users/verifyotp", func(c *gin.Context) { VerifyOTPHandler(c, queries) })

	r.Run(":8080")
}
