package handler

import (
	"dnuf/helper"
	"dnuf/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService: userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.WrapperResponse(http.StatusBadRequest, false, err.Error(), ""))
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.WrapperResponse(http.StatusInternalServerError, false, "Internal Server Error", ""))
		return
	}

	formated := user.FormatUser(newUser, "token")

	response := helper.WrapperResponse(http.StatusOK, true, "Register success", formated)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.WrapperResponse(http.StatusBadRequest, false, err.Error(), ""))
		return
	}

	rsLogin, err := h.userService.Login(input)
	if err != nil {
		c.JSON(http.StatusUnauthorized, helper.WrapperResponse(http.StatusUnauthorized, false, err.Error(), ""))
		return
	}

	formated := user.FormatUser(rsLogin, "token")

	c.JSON(http.StatusOK, helper.WrapperResponse(http.StatusOK, true, "Login success", formated))
}

func (h *userHandler) CheckEmail(c *gin.Context) {
	var input user.CheckEmailInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.WrapperResponse(http.StatusBadRequest, false, err.Error(), ""))
		return
	}

	errEmail := h.userService.CheckEmail(input)
	if errEmail != nil {
		c.JSON(http.StatusConflict, helper.WrapperResponse(http.StatusConflict, false, errEmail.Error(), ""))
		return
	}

	c.JSON(http.StatusOK, helper.WrapperResponse(http.StatusOK, true, "Email is available", ""))
}
