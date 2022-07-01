package campaign

type CampaignFormatter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	Description      string `json:"description"`
	Perks            string `json:"perks"`
	BackerCount      int    `json:"backer_count"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
	Image            string `json:"image"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	var image string
	if len(campaign.CampaignImages) > 0 {
		image = campaign.CampaignImages[0].FileName
	}
	formatted := CampaignFormatter{
		ID:               campaign.ID,
		UserID:           campaign.UserID,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		Description:      campaign.Description,
		Perks:            campaign.Perks,
		BackerCount:      campaign.BackerCount,
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		Slug:             campaign.Slug,
		Image:            image,
	}
	return formatted
}
