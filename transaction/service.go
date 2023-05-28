package transaction

import (
	"course-bwastartup-backend/campaign"
	"course-bwastartup-backend/payment"
	"errors"
)

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
	paymentService     payment.Service
}

type Service interface {
	GetTransactionByCampaignId(input GetCampaignTransactionInput) ([]Transaction, error)
	GetTransactionByUserId(userId int) ([]Transaction, error)
	CreateTransaction(input CreateTransactionInput) (Transaction, error)
}

func NewService(repository Repository, campaignRepository campaign.Repository, paymentService payment.Service) *service {
	return &service{repository, campaignRepository, paymentService}
}

// func NewService(repository Repository, campaignRepository campaign.Repository) *service {
// 	return &service{repository, campaignRepository}
// }

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

func (s *service) CreateTransaction(input CreateTransactionInput) (Transaction, error) {
	transaction := Transaction{}
	transaction.CampaignId = input.CampaignId
	transaction.Amount = input.Amount
	transaction.UserId = input.User.ID
	transaction.Status = "pending"

	// codeCandidate := fmt.Sprintf("ORDER %d000", input.User.ID)
	// transaction.Code = slug.Make(codeCandidate)
	// transaction.Code = ""

	newTransactions, err := s.repository.Save(transaction)
	if err != nil {
		return newTransactions, err
	}

	paymentTransaction := payment.Transaction{
		Id:     newTransactions.Id,
		Amount: newTransactions.Amount,
	}

	paymentURL, err := s.paymentService.GetPaymentURL(paymentTransaction, input.User)
	if err != nil {
		return newTransactions, err
	}

	newTransactions.PaymentURL = paymentURL

	newTransactions, err = s.repository.Update(newTransactions)
	if err != nil {
		return newTransactions, err
	}

	return newTransactions, nil
}
