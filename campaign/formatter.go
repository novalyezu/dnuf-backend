package campaign

import (
	"strings"
)

type CampaignFormatter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
	Image            string `json:"image"`
}

type CampaignDetailFormatter struct {
	ID               int                      `json:"id"`
	UserID           int                      `json:"user_id"`
	Name             string                   `json:"name"`
	ShortDescription string                   `json:"short_description"`
	Description      string                   `json:"description"`
	Perks            []string                 `json:"perks"`
	BackerCount      int                      `json:"backer_count"`
	GoalAmount       int                      `json:"goal_amount"`
	CurrentAmount    int                      `json:"current_amount"`
	Slug             string                   `json:"slug"`
	Images           []CampaignImageFormatter `json:"images"`
	User             CampaignUserFormatter    `json:"user"`
}

type CampaignUserFormatter struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Occupation     string `json:"occupation"`
	AvatarFileName string `json:"avatar_file_name"`
}

type CampaignImageFormatter struct {
	ID        int    `json:"id"`
	FileName  string `json:"file_name"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	formatted := CampaignFormatter{}
	formatted.ID = campaign.ID
	formatted.UserID = campaign.UserID
	formatted.Name = campaign.Name
	formatted.ShortDescription = campaign.ShortDescription
	formatted.GoalAmount = campaign.GoalAmount
	formatted.CurrentAmount = campaign.CurrentAmount
	formatted.Slug = campaign.Slug
	formatted.Image = ""

	if len(campaign.CampaignImages) > 0 {
		formatted.Image = campaign.CampaignImages[0].FileName
	}

	return formatted
}

func FormatCampaigns(campaign []Campaign) []CampaignFormatter {
	formatted := []CampaignFormatter{}

	for _, item := range campaign {
		formattedItem := FormatCampaign(item)
		formatted = append(formatted, formattedItem)
	}

	return formatted
}

func FormatCampaignDetail(campaign Campaign) CampaignDetailFormatter {
	formatted := CampaignDetailFormatter{}
	formatted.ID = campaign.ID
	formatted.UserID = campaign.UserID
	formatted.Name = campaign.Name
	formatted.ShortDescription = campaign.ShortDescription
	formatted.Description = campaign.Description
	formatted.BackerCount = campaign.BackerCount
	formatted.GoalAmount = campaign.GoalAmount
	formatted.CurrentAmount = campaign.CurrentAmount
	formatted.Slug = campaign.Slug
	formatted.Images = []CampaignImageFormatter{}
	formatted.User = CampaignUserFormatter{
		ID:             campaign.User.ID,
		Name:           campaign.User.Name,
		Occupation:     campaign.User.Occupation,
		AvatarFileName: campaign.User.AvatarFileName,
	}

	formatted.Perks = append(formatted.Perks, strings.Split(campaign.Perks, "|")...)

	for _, item := range campaign.CampaignImages {
		formattedItem := CampaignImageFormatter{
			ID:        item.ID,
			FileName:  item.FileName,
			IsPrimary: item.IsPrimary == 1,
		}
		formatted.Images = append(formatted.Images, formattedItem)
	}

	return formatted
}
