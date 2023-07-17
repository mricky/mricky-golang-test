package handler

import (
	"mricky-golang-test/helper"
	"mricky-golang-test/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service user.UserService
}


func ImplUserHandler(service user.UserService) *userHandler {
	return &userHandler{service }
}

func (h *userHandler) RegisterUser(c *gin.Context){
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatError(err)

		errorMessage := gin.H{"errors": errors} // interface

		response := helper.APIResponse("Register account failed",http.StatusUnprocessableEntity,"error",errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.service.RegisterUser(input)

	
	if err != nil {
		response := helper.APIResponse("Register account failed",http.StatusBadRequest,"error",nil)
		c.JSON(http.StatusOK, response)
		return
	}
	
	formater := user.FormaterUser(newUser,"jwtoken")
	response := helper.APIResponse("Account has been registered",http.StatusOK,"success",formater)
	
	c.JSON(http.StatusOK, response)
}