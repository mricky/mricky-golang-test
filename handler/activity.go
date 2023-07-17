package handler

import (
	"mricky-golang-test/activity"
	"mricky-golang-test/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type activityHandler struct {
	service activity.ActivityService
}

func ImplActivityHandler(service activity.ActivityService) *activityHandler {
	return &activityHandler{service}
}

// api/v1/activities?skill_id=1&user_id

func (h *activityHandler) GetActivities(c *gin.Context){
	// token, _ := strconv.Atoi(c.Query("token")) // untuk token query
	// currentUser := c.MushGet("currentUser")
	activities, err := h.service.GetActivities()
	if err != nil {
		response := helper.APIResponse("Error to get activies",http.StatusBadRequest,"error",nil)
		c.JSON(http.StatusBadRequest,response)
		return
	}

	response := helper.APIResponse("List of Activitiess",http.StatusAccepted,"success",activities)
	c.JSON(http.StatusAccepted,response)
}

func (h *activityHandler) SaveActivity(c *gin.Context){
	
	var input activity.CreateActivityInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors} // interface

		response := helper.APIResponse("Create activity failed",http.StatusUnprocessableEntity,"error",errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newActivity, err := h.service.Create(input)
	// check skill related with user skill
	// check userid
	// validate greater

	response := helper.APIResponse("Activity has been created",http.StatusOK,"success",newActivity)
	
	c.JSON(http.StatusOK, response)
}