package users

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

	users := &Module{db: db}

	router.GET("/users", users.list())
	router.POST("/users", users.create())
	router.PUT("/users/:id", users.update())
}
