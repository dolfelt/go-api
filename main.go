package main

import (
	"log"

	"github.com/dolfelt/go-api/data"
	"github.com/dolfelt/go-api/modules/auth"
	"github.com/dolfelt/go-api/modules/items"
	"github.com/dolfelt/go-api/modules/users"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func connectDB() *data.DB {
	// Open our database connection.
	db, err := gorm.Open("postgres", "user=postgres dbname=api sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	return &data.DB{DB: db}
}

func runMigration(db *data.DB) {
	db.AutoMigrate(
		&data.User{},
		&data.Item{},
	)
}

func main() {
	binding.Validator = new(defaultValidator)
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	db := connectDB()
	defer db.Close()

	runMigration(db)

	auth.Register(router, db)

	authRoute := router.Group("/", data.JWTAuthMiddleware())
	{
		auth.RegisterAuth(authRoute, db)
		users.Register(authRoute, db)
		items.Register(authRoute, db)
	}
	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()
	// router.Run(":3000") for a hard coded port
}
