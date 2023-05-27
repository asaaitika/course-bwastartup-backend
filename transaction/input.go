package transaction

import "course-bwastartup-backend/user"

type GetCampaignTransactionInput struct {
	Id   int `uri:"id" binding:"required"`
	User user.User
}

type CreateTransactionInput struct {
	Amount     int `uri:"amount" binding:"required"`
	CampaignId int `uri:"campaign_id" binding:"required"`
	User       user.User
}
