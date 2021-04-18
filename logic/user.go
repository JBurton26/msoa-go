package logic

import (
	"context"

	user "github.com/JBurton26/msoa-go/protos/user"
	tools "github.com/JBurton26/msoa-go/tools"
	"google.golang.org/api/iterator"

	hclog "github.com/hashicorp/go-hclog"
)

//User something
type User struct {
	log  hclog.Logger
	path string
}

// NewUser Initialise
func NewUser(l hclog.Logger, s string) *User {
	return &User{l, s}
}

// GetUserByID DEPRECATED METHOD
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
	u.log.Info("Handle GetUserByUN", "username", ur.GetUsername())

	app, err := tools.SetFirebase(ctx, u.log, u.path)
	client, err := app.Firestore(context.Background())
	if err != nil {
		u.log.Info("Error: Declaring Firestore app", "err", err.Error())
	}
	defer client.Close()

	iter := client.Collection("users").Where("Username", "==", ur.GetUsername()).Documents(ctx)

	doc, err := iter.Next()
	if err == iterator.Done {
		u.log.Info("Handle GetUserByUN: User Does Not Exist", "username", ur.GetUsername())
		u1 := user.UserResponse{AccessLevel: user.AccessLevel(1), ID: "NO"}
		u.log.Info("Handle GetUserByUN", "Response", u1.String())
		return &u1, nil
	}
	if err != nil {
		u.log.Info("Handle GetUserByUN", "Error", err)
		u1 := user.UserResponse{AccessLevel: user.AccessLevel(1), ID: "NO"}
		u.log.Info("Handle GetUserByUN", "Response", u1.String())
		return &u1, nil
	}
	dets := doc.Data()
	accs := user.AccessLevel(dets["AccessLevel"].(int64))
	u.log.Info("Handle GetUserByUN", "AccessLevel", accs)
	u1 := user.UserResponse{
		Username:    dets["Username"].(string),
		Name:        dets["Name"].(string),
		ID:          dets["ID"].(string),
		Location:    dets["Location"].(string),
		AccessLevel: accs,
	}
	u.log.Info("Handle GetUserByUN", "User", u1.String())
	return &u1, nil
}

// GetAccess blahblablah
func (u *User) GetAccess(ctx context.Context, ur *user.IDUserRequest) (*user.AccessResponse, error) {
	access := user.AccessResponse{
		AccessLevel: 3,
	}
	u.log.Info("Handle GetAccess", "id", ur.GetID(), "New Access", access.GetAccessLevel())
	return &access, nil
}

// Login blahblablah
func (u *User) Login(ctx context.Context, ur *user.AuthsRequest) (*user.AccessResponse, error) {
	u.log.Info("Handle Login", "username", ur.GetUsername())

	app, err := tools.SetFirebase(ctx, u.log, u.path)
	client, err := app.Firestore(context.Background())
	if err != nil {
		u.log.Info("Error: Declaring Firestore app", "err", err.Error())
	}
	defer client.Close()

	iter := client.Collection("users").Where("Username", "==", ur.GetUsername()).Where("Pass", "==", ur.GetPass()).Documents(ctx)
	doc, err := iter.Next()
	if err == iterator.Done {
		u.log.Info("Handle Login: Username or password incorrect", "username", ur.GetUsername())
		u1 := user.AccessResponse{AccessLevel: user.AccessLevel(1)}
		u.log.Info("Handle Login", "Response", u1.String())
		return &u1, nil
	}
	if err != nil {
		u.log.Info("Handle Login", "Error", err)
		u1 := user.AccessResponse{AccessLevel: user.AccessLevel(1)}
		u.log.Info("Handle Login", "Response", u1.String())
		return &u1, nil
	}
	ac := doc.Data()["AccessLevel"].(int64)
	lo := doc.Data()["Location"].(string)
	access := user.AccessResponse{AccessLevel: user.AccessLevel(ac), Location: lo}
	return &access, nil
}
