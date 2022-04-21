package handlers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"golangEx/models"
	"golangEx/repository"
	"log"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context) {
	email := c.DefaultPostForm("email", "")       //hi.com12345
	username := c.DefaultPostForm("username", "") // ad
	password := c.DefaultPostForm("password", "") //admin
	var u = models.User{}
	u.Prepare()
	hashPass, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	u = models.User{
		Role:     false,
		Name:     username,
		Email:    email,
		Password: hashPass,
	}
	err := u.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	} else {
		repo := repository.UsersCRUD{}
		func(uc repository.UserRepository) {
			_, err := uc.Save(u)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			} else {
				c.JSON(http.StatusOK, gin.H{"message": "success"})
			}
		}(&repo)
	}
}
func FindAllUser(c *gin.Context) {
	repo := repository.UsersCRUD{}
	func(uc repository.UserRepository) {
		us := []models.User{}
		us, err := uc.FindAll()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"data": us})
		}
	}(&repo)
}

func FindUserById(c *gin.Context) {
	i := c.DefaultQuery("id", "")
	id, _ := strconv.Atoi(i)
	repo := repository.UsersCRUD{}
	func(uc repository.UserRepository) {
		u := models.User{}
		u, err := uc.FindById(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"data": u})
		}
	}(&repo)
}

func FindUserByEmail(c *gin.Context) {
	email := c.DefaultPostForm("email", "")
	repo := repository.UsersCRUD{}
	func(uc repository.UserRepository) {
		u := models.User{}
		u, err := uc.FindByEmail(email)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"data": u})
		}
	}(&repo)
}
func UpDateUserById(c *gin.Context) {
	i := c.DefaultQuery("id", "")
	id, _ := strconv.Atoi(i)
	repo := repository.UsersCRUD{}
	func(uc repository.UserRepository) {
		u := models.User{}
		u, err := uc.FindById(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		} else {
			u.Name = c.DefaultPostForm("username", "")
			password := c.DefaultPostForm("password", "")
			u.Password, _ = bcrypt.GenerateFromPassword([]byte(password), 14)
			_, err := uc.Update(id, u)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			} else {
				c.JSON(http.StatusOK, gin.H{"data": u})
			}
		}
	}(&repo)
}
func DeleteUserById(c *gin.Context) {
	i := c.DefaultQuery("id", "")
	id, _ := strconv.Atoi(i)
	repo := repository.UsersCRUD{}
	func(uc repository.UserRepository) {
		log.Println(id)
		stt, err := uc.Delete(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"data": stt})
		}
	}(&repo)
}
