package data

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// CreateJWT will create a token
func CreateJWT(userID uint) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": userID,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return tokenString
}

func parseJWT(tokenString string) jwt.MapClaims {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		fmt.Println(err)
		return nil
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims
	}

	return nil
}

// JWTAuthMiddleware middleware
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		success := parseJWT(token)
		if success == nil {
			c.JSON(401, &gin.H{
				"error": "The token you are using is invalid or expired",
			})
			c.AbortWithStatus(401)
		}

		c.Set("token:claims", success)
		c.Set("token", token)
		c.Set("userID", success["user"])

		c.Next()
	}
}

// GetAuthUser returns the user who authorized the request
func GetAuthUser(db *DB, c *gin.Context) User {
	var user User

	db.First(&user, c.MustGet("userID"))

	return user
}
