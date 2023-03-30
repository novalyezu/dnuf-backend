package handler

import (
	"dnuf/campaign"
	"dnuf/helper"
	"dnuf/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	campaignService campaign.Service
}

type ErrorType struct {
	Status  int
	Message string
}

func NewCampaignHandler(campaignService campaign.Service) *campaignHandler {
	return &campaignHandler{campaignService: campaignService}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	qUserID := c.Query("user_id")
	userID := 0
	if qUserID != "" {
		cUserID, err := strconv.Atoi(c.Query("user_id"))
		if err == nil {
			userID = cUserID
		}
	}

	rsCampaigns, err := h.campaignService.GetCampaigns(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.WrapperResponse(http.StatusInternalServerError, false, "Internal Server Error", ""))
		return
	}

	formatted := campaign.FormatCampaigns(rsCampaigns)

	response := helper.WrapperResponse(http.StatusOK, true, "Get campaigns success", formatted)

	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) GetCampaignBySlug(c *gin.Context) {
	slug := c.Param("slug")

	rsCampaign, err := h.campaignService.GetCampaignBySlug(slug)
	if err != nil {
		errType := ErrorType{
			Status:  http.StatusInternalServerError,
			Message: "Internal Server Error",
		}
		if err.Error() == "NOT_FOUND" {
			errType.Status = http.StatusNotFound
			errType.Message = "Not Found"
		}
		c.JSON(errType.Status, helper.WrapperResponse(errType.Status, false, errType.Message, ""))
		return
	}

	formatted := campaign.FormatCampaignDetail(rsCampaign)

	response := helper.WrapperResponse(http.StatusOK, true, "Get campaign success", formatted)

	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) GetCampaignByID(c *gin.Context) {
	var campaignID = c.Param("campaignID")

	rsCampaign, err := h.campaignService.GetCampaignByID(campaignID)
	if err != nil {
		errType := ErrorType{
			Status:  http.StatusInternalServerError,
			Message: "Internal Server Error",
		}
		if err.Error() == "NOT_FOUND" {
			errType.Status = http.StatusNotFound
			errType.Message = "Not Found"
		}
		c.JSON(errType.Status, helper.WrapperResponse(errType.Status, false, errType.Message, ""))
		return
	}

	formatted := campaign.FormatCampaignDetail(rsCampaign)

	response := helper.WrapperResponse(http.StatusOK, true, "Get campaign success", formatted)

	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) CreateCampaign(c *gin.Context) {
	var input campaign.CreateCampaignInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.WrapperResponse(http.StatusBadRequest, false, err.Error(), ""))
		return
	}
	currUser := c.MustGet("currentUser").(user.User)
	input.User = currUser

	newCampaign, errCreateCampaign := h.campaignService.CreateCampaign(input)
	if errCreateCampaign != nil {
		c.JSON(http.StatusInternalServerError, helper.WrapperResponse(http.StatusInternalServerError, false, "Internal Server Error", ""))
		return
	}

	formatted := campaign.FormatCampaignDetail(newCampaign)

	response := helper.WrapperResponse(http.StatusOK, true, "Create campaign success", formatted)

	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) UpdateCampaign(c *gin.Context) {
	var campaignID = c.Param("campaignID")
	var input campaign.CreateCampaignInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.WrapperResponse(http.StatusBadRequest, false, err.Error(), ""))
		return
	}
	currUser := c.MustGet("currentUser").(user.User)
	input.User = currUser

	updateCampaign, errUpdateCampaign := h.campaignService.UpdateCampaign(campaignID, input)
	if errUpdateCampaign != nil {
		errType := ErrorType{
			Status:  http.StatusInternalServerError,
			Message: "Internal Server Error",
		}
		if errUpdateCampaign.Error() == "NOT_FOUND" {
			errType.Status = http.StatusNotFound
			errType.Message = "Not Found"
		}
		c.JSON(errType.Status, helper.WrapperResponse(errType.Status, false, errType.Message, ""))
		return
	}

	formatted := campaign.FormatCampaignDetail(updateCampaign)

	response := helper.WrapperResponse(http.StatusOK, true, "Update campaign success", formatted)

	c.JSON(http.StatusOK, response)
}
