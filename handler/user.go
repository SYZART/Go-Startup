package handler

import (
	"bwastartup/auth"
	"bwastartup/helper"
	"bwastartup/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}

		respponse := helper.APIResponse("Registration Failed", http.StatusUnprocessableEntity, "error", errorMessage)

		c.JSON(http.StatusBadRequest, respponse)
		return
	}
	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		respponse := helper.APIResponse("Registration Failed", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, respponse)
		return
	}

	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {
		respponse := helper.APIResponse("Generate Token Invalid", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, respponse)
		return
	}

	formatter := user.UserFormatJSON(newUser, token)
	respponse := helper.APIResponse("Account has been created", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, respponse)
}

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput
	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}
		respponse := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, respponse)
		return
	}
	loggedinUser, err := h.userService.Login(input)

	if err != nil {

		errorMessage := gin.H{"errors": err.Error()}
		respponse := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, respponse)
		return

	}
	tokenUserLogin, err := h.authService.GenerateToken(loggedinUser.ID)
	if err != nil {
		respponse := helper.APIResponse("Generate Token Invalid", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, respponse)
		return
	}
	formatter := user.UserFormatJSON(loggedinUser, tokenUserLogin)
	respponse := helper.APIResponse("Successfully Loggin", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, respponse)
}

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	var input user.CheckEmailInput
	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}
		respponse := helper.APIResponse("Email has already", http.StatusUnprocessableEntity, "error", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, respponse)
		return
	}
	IsEmailAvailable, err := h.userService.IsEmailAvailabe(input)
	if err != nil {

		errorMessage := gin.H{"errors": "Server Error"}
		respponse := helper.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, respponse)
		return

	}
	data := gin.H{
		"is_available": IsEmailAvailable,
	}
	metaMessege := "Email address has been registered"
	if IsEmailAvailable {
		metaMessege = "Email has been available"
	}
	response := helper.APIResponse(metaMessege, http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)

}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	userID := 23

	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Avatar successfully uploaded", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)

}
