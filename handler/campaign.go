package handler

import (
	"dnuf/campaign"
	"dnuf/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	campaignService campaign.Service
}

func NewCampaignHandler(campaignService campaign.Service) *campaignHandler {
	return &campaignHandler{campaignService: campaignService}
}

func (h *campaignHandler) GetAll(c *gin.Context) {
	rsCampaigns, err := h.campaignService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.WrapperResponse(http.StatusInternalServerError, false, "Internal Server Error", ""))
		return
	}

	var formatted []campaign.CampaignFormatter
	for _, item := range rsCampaigns {
		formattedItem := campaign.FormatCampaign(item)
		formatted = append(formatted, formattedItem)
	}

	response := helper.WrapperResponse(http.StatusOK, true, "Register success", formatted)

	c.JSON(http.StatusOK, response)
}
