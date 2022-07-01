package handler

import (
	"dnuf/campaign"
	"dnuf/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	campaignService campaign.Service
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

func (h *campaignHandler) GetCampaign(c *gin.Context) {
	slug := c.Param("slug")

	rsCampaign, err := h.campaignService.GetCampaign(slug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.WrapperResponse(http.StatusInternalServerError, false, "Internal Server Error", ""))
	}

	formatted := campaign.FormatCampaignDetail(rsCampaign)

	response := helper.WrapperResponse(http.StatusOK, true, "Get campaign success", formatted)

	c.JSON(http.StatusOK, response)
}
