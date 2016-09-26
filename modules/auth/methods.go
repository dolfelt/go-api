package auth

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/dolfelt/go-api/data"
	"github.com/gin-gonic/gin"
)

type loginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func loginError(c *gin.Context) {
	c.JSON(401, &gin.H{
		"error": "Email or password is incorrect",
	})
	c.Abort()
}

func (m *Module) login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input loginInput
		c.Bind(&input)

		var user data.User
		m.db.Where("email = ?", input.Email).First(&user)

		fmt.Println(user.ID)
		if user.ID == 0 {
			loginError(c)
			return
		}

		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
		if err != nil {
			fmt.Println(err)
			loginError(c)
			return
		}
		token := data.CreateJWT(user.ID)

		c.JSON(200, gin.H{
			"status": "success",
			"user":   user,
			"token":  token,
		})
	}
}

func (m *Module) getLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.MustGet("userID")
		var user data.User
		m.db.First(&user, userID)

		c.JSON(200, gin.H{
			"status": "replay",
			"user":   user,
			"token":  c.MustGet("token").(string),
		})
	}
}
