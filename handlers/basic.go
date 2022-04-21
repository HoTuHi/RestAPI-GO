package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"golangEx/connections"
	"golangEx/models"
	"log"
	"net/http"
	"time"
)

const MySecretKey = "hotuhi"

// ROUTERS
func Register(c *gin.Context) {
	//var data map[string]string
	email := c.DefaultPostForm("email", "")
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")

	if username != "" {
		if password != "" {
			hashPass, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
			userTest := models.User{
				Role:     false,
				Name:     username,
				Email:    email,
				Password: hashPass,
				CreateAt: time.Now(),
			}
			connections.DB.Create(&userTest)
			c.JSON(http.StatusOK, gin.H{"message": "success"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid"})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid"})
	}
}
func Login(c *gin.Context) {
	//get user and password from body
	email := c.DefaultPostForm("email", "")
	password := c.DefaultPostForm("password", "")
	var user = models.User{}
	connections.DB.Where("email = ?", email).First(&user)
	// if found user in database
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "NotFound"})
	} else
	//  compare pass in database, if wrong password
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(password)); err != nil {
		log.Println("Wrong Username/Password")
		c.JSON(http.StatusBadRequest, gin.H{"message": "Wrong Username/Password"})
	} else {
		// claims have Issuer, ExpiresAt
		claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 10).Unix(),
			Issuer:    user.Email,
			Subject:   "",
		})
		token, error := claims.SignedString([]byte(MySecretKey))
		if error != nil {
			log.Println(error.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Cannot login"})
		}
		//cookie, _ := c.Cookie("jwt")
		c.SetCookie("jwt", string(token), 3600, "/", "localhost", false, true)
		//c.SetCookie("user", user.Email, 60, "/", "localhost", false, true)
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	}

}
func Logout(c *gin.Context) {
	c.Cookie("jwt")
	c.SetCookie("jwt", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusUnauthorized, gin.H{"message": "success"})
}
func UserInfo(c *gin.Context) {
	cookie, _ := c.Cookie("jwt")
	log.Println(cookie)
	token, _ := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(MySecretKey), nil
	})
	claims := token.Claims.(*jwt.StandardClaims)

	var user = models.User{}
	connections.DB.Where("id = ?", claims.Issuer).First(&user)
	c.JSON(http.StatusOK, user)
}
func UpdateUser(c *gin.Context) {
	//var data map[string]string
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")
	hashPass, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	cookie, _ := c.Cookie("jwt")
	log.Println(cookie)
	token, _ := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(MySecretKey), nil
	})
	claims := token.Claims.(*jwt.StandardClaims)
	var bufferUser = models.User{}
	connections.DB.Where("email=?", claims.Issuer).First(&bufferUser)
	if err := bcrypt.CompareHashAndPassword(bufferUser.Password, []byte(password)); err != nil {
		bufferUser.Password = hashPass
		bufferUser.Name = username
		connections.DB.Save(bufferUser)
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Not Change!"})

	}
}
func DeleteUser(c *gin.Context) {
	if CheckAdmin(c) {
		email := c.DefaultPostForm("email", "")
		var fUser = models.User{}
		result := connections.DB.Where("email=?", email).First(&fUser)
		//log.Println(result.Error)
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "NotFound"})
		} else {
			connections.DB.Delete(models.User{}, "email=?", email)
			c.JSON(http.StatusOK, gin.H{"message": "Success"})
		}

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unaccepted!"})
	}
}
func GetAllUser(c *gin.Context) {
	if CheckAdmin(c) {
		var users []models.User
		connections.DB.Find(&users)
		c.JSON(http.StatusOK, users)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unaccepted!"})
	}

}
func CheckAdmin(c *gin.Context) bool {
	cookie, _ := c.Cookie("jwt")
	token, _ := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(MySecretKey), nil
	})
	claims := token.Claims.(*jwt.StandardClaims)
	var bufferUser = models.User{}
	connections.DB.Where("email=?", claims.Issuer).First(&bufferUser)
	if bufferUser.Role {
		return true
	}
	return false
}
