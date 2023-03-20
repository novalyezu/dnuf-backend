package campaign

import (
	"strconv"

	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaigns(userID int) ([]Campaign, error)
	GetCampaign(slug string) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetCampaigns(userID int) ([]Campaign, error) {
	if userID != 0 {
		rsCampaigns, err := s.repository.FindByUserID(userID)
		if err != nil {
			return rsCampaigns, err
		}
		return rsCampaigns, nil
	}
	rsCampaigns, err := s.repository.FindAll()
	if err != nil {
		return rsCampaigns, err
	}
	return rsCampaigns, nil
}

func (s *service) GetCampaign(slug string) (Campaign, error) {
	rsCampaign, err := s.repository.FindBySlug(slug)
	if err != nil {
		return rsCampaign, err
	}
	return rsCampaign, nil
}

func (s *service) CreateCampaign(input CreateCampaignInput) (Campaign, error) {
	campaign := Campaign{}
	campaign.Name = input.Name
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.GoalAmount = input.GoalAmount
	campaign.Perks = input.Perks
	campaign.UserID = input.User.ID
	campaign.Slug = slug.Make(input.Name + strconv.Itoa(input.User.ID))

	newCampaign, err := s.repository.Save(campaign)
	if err != nil {
		return newCampaign, err
	}
	return newCampaign, nil
}
