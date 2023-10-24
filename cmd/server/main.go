package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tonytangdev/go-rest-api/internal"
	"github.com/tonytangdev/go-rest-api/internal/handler"
	"github.com/tonytangdev/go-rest-api/internal/middleware"
)

func main() {
	db, err := internal.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = internal.InitializeDB(db)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	r := gin.Default()
	r.Use(middleware.DBMiddleware(db)) // Register the middleware

	r.GET("/users", handler.GetUsers)
	r.POST("/users", handler.PostUser)
	r.DELETE("/users/:id", handler.DeleteUser)
	r.PUT("/users/:id", handler.UpdateUser)

	r.Run() // listen and serve on 0.0.0.0:8080
}
