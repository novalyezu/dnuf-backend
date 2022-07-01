package campaign

type Service interface {
	GetCampaigns(userID int) ([]Campaign, error)
	GetCampaign(slug string) (Campaign, error)
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
