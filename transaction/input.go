package transaction

import "bwastartup/user"

type GetTranscationsCampaignInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}
