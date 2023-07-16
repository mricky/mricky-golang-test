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
	activities, err := h.service.GetActivities()
	if err != nil {
		response := helper.APIResponse("Error to get activies",http.StatusBadRequest,"error",nil)
		c.JSON(http.StatusBadRequest,response)
		return
	}

	response := helper.APIResponse("List of Activitiess",http.StatusAccepted,"success",activities)
	c.JSON(http.StatusAccepted,response)
}