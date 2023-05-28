package transaction

import "course-bwastartup-backend/user"

type GetCampaignTransactionInput struct {
	Id   int `uri:"id" binding:"required"`
	User user.User
}

type CreateTransactionInput struct {
	Amount     int `json:"amount" binding:"required"`
	CampaignId int `json:"campaign_id" binding:"required"`
	User       user.User
}

type TransactionNotificationInput struct {
	TransactionStatus string `json:"transaction_amount"`
	OrderID           string `json:"order_id" `
	PaymentType       string `json:"payment_type"`
	FraudStatus       string `json:"fraud_status"`
}
