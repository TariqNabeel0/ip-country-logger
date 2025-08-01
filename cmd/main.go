package main

import (
	"ip-country-logger/config"
	"ip-country-logger/handlers"
	"ip-country-logger/models"

	"github.com/gin-gonic/gin"
)

func main (){
	r := gin.Default()
config.ConnectDB()

config.DB.AutoMigrate(&models.Visit{})

r.POST("/visit", handlers.PostVisit)
r.GET("/visits", handlers.GetVisit)
r.GET("/summary", handlers.GetSummary)

r.Run(":8080")

}