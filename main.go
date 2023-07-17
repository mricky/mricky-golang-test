package main

import (
	"fmt"
	"log"
	"mricky-golang-test/activity"
	"mricky-golang-test/auth"
	"mricky-golang-test/handler"
	"mricky-golang-test/helper"
	"mricky-golang-test/skill"
	"mricky-golang-test/user"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
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
	authService := auth.ImplService()
	userHandler := handler.ImplUserHandler(userService,authService)
	
	
	activityRepository := activity.ImplActivityRepository(db)
    activityService := activity.ImplActivityService(activityRepository)
    activityHandler := handler.ImplActivityHandler(activityService)
	
	skillRepository := skill.ImplSkillRepository(db)
	skillService := skill.ImplSkillService(skillRepository)
	skillHandler := handler.ImplSkillHandler(skillService)
	

	// test jwt
	
	fmt.Println(authService.GenerateToken(1))
	
	router := gin.Default()
	api := router.Group("/v1")

	api.POST("/users",authMiddleware(authService, userService),userHandler.RegisterUser)
	api.POST("auth/login",userHandler.Login)

	api.POST("/activities",authMiddleware(authService, userService),activityHandler.SaveActivity)
	api.GET("/activities",authMiddleware(authService, userService),activityHandler.GetActivities)
	
	// SKILL
	api.GET("/skills",authMiddleware(authService, userService),skillHandler.GetSkills)
	
	router.Run()
	
}

func authMiddleware(authService auth.Service, userService user.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		inputToken := c.Query("token")

		if len(strings.TrimSpace(inputToken)) == 0 { 
			response := helper.APIResponse("Unauthorize",http.StatusUnauthorized,"error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		} 

		token, err := authService.ValidateToken(inputToken)

		if err != nil {
			response := helper.APIResponse("Unauthorize",http.StatusUnauthorized,"error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorize",http.StatusUnauthorized,"error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		
		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)

		if err != nil{
			response := helper.APIResponse("Unauthorize",http.StatusUnauthorized,"error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser",user)
	}
}