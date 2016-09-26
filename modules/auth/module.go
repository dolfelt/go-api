package auth

import (
	"github.com/dolfelt/go-api/data"
	"github.com/gin-gonic/gin"
)

// Module exports all methods
type Module struct {
	db *data.DB
}

// Register creates the routes
func Register(router gin.IRouter, db *data.DB) {

	authModule := &Module{db: db}

	router.POST("/login", authModule.login())
}

// RegisterAuth creates routes
func RegisterAuth(router gin.IRouter, db *data.DB) {
	authModule := &Module{db: db}

	router.GET("/login", authModule.getLogin())
	router.GET("/users/me", authModule.getLogin())
}
