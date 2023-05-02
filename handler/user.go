package handler

import (
	"course-bwastartup-backend/helper"
	"course-bwastartup-backend/user"
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
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// token, err := h.jwtService.GenerateToken()
	formatter := user.FormatUser(newUser, "token123")

	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}

func (h *userHandler) Login(c *gin.Context) {
	/*1. user memasukkan input (email & password)
	2. input ditangkap handler
	3. mapping dari input user ke input struct
	4. input struct passing ke service
	5. di dalam service, mencari dgn bantuan repository user dengan email
	6. mencocokan password*/

}
