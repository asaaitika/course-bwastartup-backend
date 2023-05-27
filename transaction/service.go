package transaction

import (
	"course-bwastartup-backend/campaign"
	"errors"
)

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
}

type Service interface {
	GetTransactionByCampaignId(input GetCampaignTransactionInput) ([]Transaction, error)
	GetTransactionByUserId(userId int) ([]Transaction, error)
}

func NewService(repository Repository, campaignRepository campaign.Repository) *service {
	return &service{repository, campaignRepository}
}

func (s *service) GetTransactionByCampaignId(input GetCampaignTransactionInput) ([]Transaction, error) {
	campaign, err := s.campaignRepository.FindByID(input.Id)
	if err != nil {
		return []Transaction{}, err
	}

	if campaign.UserID != input.User.ID {
		return []Transaction{}, errors.New("not an owner of the campaign")
	}

	transaction, err := s.repository.GetByCampaignId(input.Id)
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (s *service) GetTransactionByUserId(userId int) ([]Transaction, error) {
	transactions, err := s.repository.GetByUserId(userId)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
