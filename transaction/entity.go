package transaction

import (
	"bwastartup/user"
	"time"
)

type Transaction struct {
	ID         int
	UserID     int
	CampaignID int
	Amount     int
	Status     string
	Code       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	User       user.User
}
