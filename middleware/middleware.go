package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golangEx/connections"
	"golangEx/handlers"
	"golangEx/models"
	"log"
	"net/http"
	"strconv"
)

func SetUserStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("jwt")
		_, err = jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(handlers.MySecretKey), nil
		})
		if err == nil {
			log.Println("is login")
			c.Set("is_logged_in", true)
		} else {
			log.Println("not login")
			c.Set("is_logged_in", false)
		}
		c.Next()
	}
}
func EnsureLogin() gin.HandlerFunc {
	log.Println("in EnsureLogin Middleware")
	return func(c *gin.Context) {
		loggedInInterface, _ := c.Get("is_logged_in")
		loggedIn := loggedInInterface.(bool)
		if !loggedIn {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.Abort()
		}

	}
}
func EnsureIsAdmin() gin.HandlerFunc {
	log.Println("in EnsureNotLogin Middleware")
	return func(c *gin.Context) {
		cookie, _ := c.Cookie("jwt")
		token, _ := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(handlers.MySecretKey), nil
		})
		claims := token.Claims.(*jwt.StandardClaims)
		var bufferUser = models.User{}
		connections.DB.Where("email=?", claims.Issuer).First(&bufferUser)
		if bufferUser.Role != true {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.Abort()
		}
	}
}
func EnsureIsOwner() gin.HandlerFunc {
	log.Println("in EnsureNotLogin Middleware")
	return func(c *gin.Context) {
		idString := c.DefaultPostForm("id", "")
		id, _ := strconv.Atoi(idString)
		cookie, _ := c.Cookie("jwt")
		token, _ := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(handlers.MySecretKey), nil
		})
		claims := token.Claims.(*jwt.StandardClaims)
		var bufferUser = models.User{}
		connections.DB.Where("email=?", claims.Issuer).First(&bufferUser)
		if uint(id) == bufferUser.ID {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.Abort()
		}
	}
}
