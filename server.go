/*
 * Copyright (c) 2021 VuongSoft
 * Author : HoTuHi
 * More : https://github.com/HoTuHi
 */

package main

import (
	"github.com/gin-gonic/gin"
	"golangEx/connections"
	"golangEx/handlers"
	"golangEx/middleware"
	"log"
)

func main() {
	connections.Connection()
	router := gin.Default()
	router.Use(middleware.SetUserStatus())
	router.POST("/register", handlers.CreateUser)
	router.POST("/login", handlers.Login)
	router.POST("/logout", handlers.Logout)
	log.Println("app run in http://localhost:8080/")
	router.Run(":8080")
}
