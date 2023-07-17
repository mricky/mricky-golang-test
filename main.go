package main

import (
	"log"
	"mricky-golang-test/activity"
	"mricky-golang-test/handler"
	"mricky-golang-test/skill"
	"mricky-golang-test/user"

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

	userRepository := user.ImplUserRepository(db)
	userService := user.ImplUserService(userRepository)
	userHandler := handler.ImplUserHandler(userService)

	activityRepository := activity.ImplActivityRepository(db)
    activityService := activity.ImplActivityService(activityRepository)
    activityHandler := handler.ImplActivityHandler(activityService)
	
	skillRepository := skill.ImplSkillRepository(db)
	skillService := skill.ImplSkillService(skillRepository)
	skillHandler := handler.ImplSkillHandler(skillService)
	
	
	router := gin.Default()
	api := router.Group("/v1")

	api.POST("/users",userHandler.RegisterUser)
	api.POST("auth/login",userHandler.Login)
	
	api.GET("/activities",activityHandler.GetActivities)

	// SKILL
	api.GET("/skills",skillHandler.GetSkills)
	
	router.Run()
	// next bikin login dan menggunakan middleware
}