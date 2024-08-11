package main

import (
	"tunwa/db"
	"tunwa/route"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB() // Initialize the database

	r := gin.Default()
	r.Use(gin.ErrorLogger())

	r.POST("/strong_password_steps", route.PasswordStrengthHandler)

	if err := r.Run(":3000"); err != nil {
		panic(err)
	}
}
