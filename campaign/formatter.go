package campaign

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
