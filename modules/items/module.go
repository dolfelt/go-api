package items

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

	items := &Module{db: db}

	router.GET("/items", items.list())
	router.POST("/items", items.create())
	router.PUT("/items/:id", items.update())
}
