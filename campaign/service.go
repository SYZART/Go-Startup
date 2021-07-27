package campaign

import (
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaigns(userID int) ([]Campaign, error)
	GetCampaignsByID(input GetCampaignDetailInput) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCampaigns(userID int) ([]Campaign, error) {
	if userID != 0 {
		campaigns, err := s.repository.FindByUserID(userID)
		if err != nil {
			return campaigns, err
		}
		return campaigns, nil
	}
	campaigns, err := s.repository.FindAll()
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

func (s *service) GetCampaignsByID(input GetCampaignDetailInput) (Campaign, error) {
	campaigns, err := s.repository.FindCampaignByID(input.ID)
	if err != nil {
		return campaigns, err

	}
	return campaigns, nil

}

func (s *service) CreateCampaign(input CreateCampaignInput) (Campaign, error) {
	campign := Campaign{}
	campign.Name = input.Name
	campign.ShortDescription = input.ShortDescription
	campign.Description = input.Description
	campign.Perks = input.Perks
	campign.GoalAmount = input.GoalAmount
	campign.UserID = input.User.ID
	createSlug := fmt.Sprintf("%s %d", input.Name, input.User.ID)
	campign.Slug = slug.Make(createSlug)
	newCampaign, err := s.repository.Save(campign)
	if err != nil {
		return newCampaign, err
	}
	return newCampaign, nil

}
