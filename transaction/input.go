package transaction

import (
	"bwastartup/user"
)

type GetTranscationsCampaignInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}
type CreateTransactionInput struct {
	Amount     int `json:"amount"`
	CampaignID int `json:"campaign_id"`
	User       user.User
}
