package transaction

import (
	"course-bwastartup-backend/campaign"
	"course-bwastartup-backend/payment"
	"errors"
	"strconv"
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
	ProcessPayment(input TransactionNotificationInput) error
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

func (s *service) ProcessPayment(input TransactionNotificationInput) error {
	transaction_id, _ := strconv.Atoi(input.OrderID)

	transaction, err := s.repository.GetById(transaction_id)
	if err != nil {
		return err
	}

	if input.PaymentType == "credit_card" && input.TransactionStatus == "capture" && input.FraudStatus == "accept" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "settlement" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "deny" || input.TransactionStatus == "expire" || input.TransactionStatus == "cancel" {
		transaction.Status = "cancelled"
	}

	updatedTransaction, err := s.repository.Update(transaction)
	if err != nil {
		return err
	}

	campaign, err := s.campaignRepository.FindByID(updatedTransaction.CampaignId)
	if err != nil {
		return err
	}

	if updatedTransaction.Status == "paid" {
		campaign.BackerCount = campaign.BackerCount + 1
		campaign.CurrentAmout = campaign.CurrentAmout + updatedTransaction.Amount

		_, err := s.campaignRepository.Update(campaign)
		if err != nil {
			return err
		}
	}

	return nil
}
