package transaction

import (
	"bwastartup/campaign"
	"errors"
)

type Service interface {
	GetTransactionsByCampaignId(input GetTranscationsCampaignInput) ([]Transaction, error)
}

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
}

func NewService(repository Repository, campaignRepository campaign.Repository) *service {
	return &service{repository, campaignRepository}
}

func (s *service) GetTransactionsByCampaignId(input GetTranscationsCampaignInput) ([]Transaction, error) {
	campaign, err := s.campaignRepository.FindCampaignByID(input.ID)
	if err != nil {
		return []Transaction{}, err
	}
	if campaign.UserID != input.User.ID {
		return []Transaction{}, errors.New("not an owner of the campaign")
	}
	transactions, err := s.repository.GetByCampaignID(input.ID)
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}
