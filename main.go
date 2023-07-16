package main

import (
	"log"
	"mricky-golang-test/activity"
	"mricky-golang-test/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main(){
	dsn := "root:Rahasia123!@tcp(127.0.0.1:8889)/db_mricky_golang?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if (err != nil) {
    	log.Fatal(err.Error())
    }

	activityRepository := activity.ImplActivityRepository(db)
    activityService := activity.ImplActivityService(activityRepository)
    activityHandler := handler.ImplActivityHandler(activityService)
	
	router := gin.Default()
	api := router.Group("/v1")

	api.GET("/activities",activityHandler.GetActivities)

	router.Run()
	// next bikin login dan menggunakan middleware
}