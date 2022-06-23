package handler

import (
	"dnuf/auth"
	"dnuf/helper"
	"dnuf/user"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService: userService, authService: authService}
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

	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.WrapperResponse(http.StatusInternalServerError, false, "Internal Server Error", ""))
		return
	}

	formated := user.FormatUser(newUser, token)

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

	token, err := h.authService.GenerateToken(rsLogin.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.WrapperResponse(http.StatusInternalServerError, false, "Internal Server Error", ""))
		return
	}

	formated := user.FormatUser(rsLogin, token)

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

func (h *userHandler) UpdateAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.WrapperResponse(http.StatusBadRequest, false, err.Error(), ""))
		return
	}

	userID := 13
	path := "images/avatars/" + strconv.Itoa(userID) + "-" + file.Filename

	errSaveFile := c.SaveUploadedFile(file, path)
	if errSaveFile != nil {
		c.JSON(http.StatusInternalServerError, helper.WrapperResponse(http.StatusInternalServerError, false, "Internal Server Error", ""))
		return
	}

	rsUpdateAvatar, err := h.userService.UpdateAvatar(userID, path)
	if err != nil {
		os.Remove(path)
		c.JSON(http.StatusInternalServerError, helper.WrapperResponse(http.StatusInternalServerError, false, "Internal Server Error", ""))
		return
	}

	formated := user.FormatUser(rsUpdateAvatar, "token")

	c.JSON(http.StatusOK, helper.WrapperResponse(http.StatusOK, true, "Update avatar success", formated))
}
