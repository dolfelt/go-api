package items

import (
	"github.com/dolfelt/go-api/data"
	"github.com/gin-gonic/gin"
)

func (m *Module) list() gin.HandlerFunc {
	return func(c *gin.Context) {
		var items []data.Item
		m.db.Find(&items)

		c.JSON(200, gin.H{
			"status": "success",
			"items":  items,
		})
	}
}

func (m *Module) create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var item data.Item
		if err := c.Bind(&item); err != nil {
			c.JSON(500, err)
			return
		}

		authUser := data.GetAuthUser(m.db, c)
		item.CreatedID = authUser.ID

		m.db.Create(&item)
		c.JSON(201, item)
	}
}

func (m *Module) update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var item data.Item
		id := c.Param("id")
		m.db.Find(&item, id)

		if err := c.Bind(&item); err != nil {
			c.JSON(500, err)
			return
		}

		m.db.Save(&item)
		c.JSON(200, item)
	}
}
