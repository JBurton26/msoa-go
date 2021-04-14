package logic

import (
	"context"

	user "github.com/JBurton26/protos/user"
	hclog "github.com/hashicorp/go-hclog"
)

//User something
type User struct {
	log hclog.Logger
}

// NewUser Initialise
func NewUser(l hclog.Logger) *User {
	return &User{l}
}

// GetUserByID blahblablah
func (u *User) GetUserByID(ctx context.Context, ur *user.IDUserRequest) (*user.UserResponse, error) {
	u1 := user.UserResponse{
		ID:          ur.GetID(),
		Name:        "Jake Burton",
		Username:    "Burty101",
		AccessLevel: 4,
	}
	u.log.Info("Handle GetUserByID", "id", ur.GetID())
	return &u1, nil
}
