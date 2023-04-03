package campaign

import (
	"errors"
	"strconv"

	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaigns(userID int) ([]Campaign, error)
	GetCampaignBySlug(slug string) (Campaign, error)
	GetCampaignByID(campaignID string) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
	UpdateCampaign(campaignID string, input CreateCampaignInput) (Campaign, error)
	SaveCampaignImage(input CreateCampaignImageInput, fileLocation string) (CampaignImage, error)
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

func (s *service) GetCampaignBySlug(slug string) (Campaign, error) {
	rsCampaign, err := s.repository.FindBySlug(slug)
	if err != nil {
		return rsCampaign, err
	}
	if rsCampaign.ID == 0 {
		return rsCampaign, errors.New("NOT_FOUND")
	}
	return rsCampaign, nil
}

func (s *service) GetCampaignByID(campaignID string) (Campaign, error) {
	rsCampaign, err := s.repository.FindByID(campaignID)
	if err != nil {
		return rsCampaign, err
	}
	if rsCampaign.ID == 0 {
		return rsCampaign, errors.New("NOT_FOUND")
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

func (s *service) UpdateCampaign(campaignID string, input CreateCampaignInput) (Campaign, error) {
	campaign, err := s.repository.FindByID(campaignID)
	if err != nil {
		return campaign, err
	}
	if campaign.ID == 0 {
		return campaign, errors.New("NOT_FOUND")
	}
	if campaign.UserID != input.User.ID {
		return campaign, errors.New("NOT_FOUND")
	}

	campaign.Name = input.Name
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.GoalAmount = input.GoalAmount
	campaign.Perks = input.Perks

	rsCampaign, errUpdate := s.repository.Update(campaign)
	if errUpdate != nil {
		return rsCampaign, errUpdate
	}
	return rsCampaign, nil
}

func (s *service) SaveCampaignImage(input CreateCampaignImageInput, fileLocation string) (CampaignImage, error) {
	isPrimary := 0
	if input.IsPrimary {
		isPrimary = 1

		err := s.repository.MarkAllImagesAsNonPrimary(input.CampaignID)
		if err != nil {
			return CampaignImage{}, err
		}
	}

	campaignImage := CampaignImage{}
	campaignImage.CampaignID = input.CampaignID
	campaignImage.IsPrimary = isPrimary
	campaignImage.FileName = fileLocation

	newCampaignImage, errCreateImage := s.repository.CreateImage(campaignImage)
	if errCreateImage != nil {
		return newCampaignImage, errCreateImage
	}
	return newCampaignImage, nil
}
