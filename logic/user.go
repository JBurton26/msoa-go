package logic

import (
	"context"

	firebase "firebase.google.com/go"
	user "github.com/JBurton26/msoa-go/protos/user"
	"google.golang.org/api/option"

	hclog "github.com/hashicorp/go-hclog"
)

//User something
type User struct {
	log hclog.Logger
}

// NewUser Initialise
func NewUser(l hclog.Logger) *User {
	sa := option.WithCredentialsFile("creds/msoa-go-user.json")

	app, err := firebase.NewApp(context.Background(), nil, sa)
	if err != nil {
		l.Info("Error", "err1", err)
	}

	client, err := app.Firestore(context.Background())
	if err != nil {
		l.Info("Error", "err2", err.Error())
	}
	defer client.Close()
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

// GetUserByUN blahblablah
func (u *User) GetUserByUN(ctx context.Context, ur *user.UNUserRequest) (*user.UserResponse, error) {
	u1 := user.UserResponse{
		ID:          ur.GetUsername(),
		Name:        "Jake Burton",
		Username:    "Burty101",
		AccessLevel: 4,
	}
	u.log.Info("Handle GetUserByUN", "username", ur.GetUsername())
	return &u1, nil
}

// GetAccess blahblablah
func (u *User) GetAccess(ctx context.Context, ur *user.IDUserRequest) (*user.AccessResponse, error) {
	access := user.AccessResponse{
		AccessLevel: 4,
	}
	u.log.Info("Handle GetAccess", "id", ur.GetID())
	return &access, nil
}

// Login blahblablah
func (u *User) Login(ctx context.Context, ur *user.AuthsRequest) (*user.AccessResponse, error) {
	userReq := user.UNUserRequest{
		Username: ur.GetUsername(),
	}
	userDet, _ := u.GetUserByUN(ctx, &userReq)
	access := user.AccessResponse{
		AccessLevel: userDet.GetAccessLevel(),
	}
	u.log.Info("Handle Login", "username", ur.GetUsername())
	return &access, nil
}
