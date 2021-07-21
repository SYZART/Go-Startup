package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}

		respponse := helper.APIResponse("Registration Failed", http.StatusBadRequest, "failed", errorMessage)

		c.JSON(http.StatusBadRequest, respponse)
		return
	}
	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		respponse := helper.APIResponse("Registration Failed", http.StatusBadRequest, "failed", nil)

		c.JSON(http.StatusBadRequest, respponse)
		return
	}
	formatter := user.UserFormatJSON(newUser, "wkwwk")
	respponse := helper.APIResponse("Account has been created", http.StatusOK, "succes", formatter)

	c.JSON(http.StatusOK, respponse)
}
