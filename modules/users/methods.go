package users

import (
	"github.com/dolfelt/go-api/data"
	"github.com/gin-gonic/gin"
)

func (m *Module) list() gin.HandlerFunc {
	return func(c *gin.Context) {
		var users []data.User
		m.db.Find(&users)

		c.JSON(200, gin.H{
			"status": "success",
			"users":  users,
		})
	}
}

func (m *Module) create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user data.User
		if err := c.Bind(&user); err != nil {
			data.PrintErrors(err, c) //c.JSON(400, err)
			return
		}
		m.db.Create(&user)
		c.JSON(201, user)
	}
}

func (m *Module) update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user data.User
		id := c.Param("id")
		m.db.Find(&user, id)

		if err := c.Bind(&user); err != nil {
			data.PrintErrors(err, c) //c.JSON(400, err)
			return
		}
		m.db.Save(&user)
		c.JSON(200, user)
	}
}
