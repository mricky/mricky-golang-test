package handler

import (
	"mricky-golang-test/helper"
	"mricky-golang-test/skill"
	"net/http"

	"github.com/gin-gonic/gin"
)

type skillHandler struct {
	service skill.SkillService
}

func ImplSkillHandler(service skill.SkillService) *skillHandler{
	return &skillHandler{service}
}

func (h *skillHandler) GetSkills(c *gin.Context){
	skills, err := h.service.GetSkills()

	if err != nil {
		response := helper.APIResponse("Error to get skill",http.StatusBadRequest,"error",nil)
		c.JSON(http.StatusBadRequest,response)
		return
	}

	response := helper.APIResponse("List of Skills",http.StatusAccepted,"success",skills)
	c.JSON(http.StatusAccepted,response)
}