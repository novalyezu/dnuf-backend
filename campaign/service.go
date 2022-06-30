package campaign

type Service interface {
	GetAll() ([]Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetAll() ([]Campaign, error) {
	rsCampaigns, err := s.repository.FindAll()
	if err != nil {
		return rsCampaigns, err
	}
	return rsCampaigns, nil
}
