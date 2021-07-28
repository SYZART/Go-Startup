package user

import (
	"time"
)

// struct ini merupakan replikasi dari field dalam database dengan nama yang singular
type User struct {
	ID             int
	Name           string
	Occupation     string
	Email          string
	PasswordHash   string
	AvatarFileName string
	Role           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
