package tools

import (
	"context"
	"math/rand"
	"time"

	firebase "firebase.google.com/go"
	hclog "github.com/hashicorp/go-hclog"
	"google.golang.org/api/option"
)

// SetFirebase Over utilised
func SetFirebase(ctx context.Context, log hclog.Logger, path string) (*firebase.App, error) {
	sa := option.WithCredentialsFile(path)
	//l.Info("Info", "sa", sa)
	app, err := firebase.NewApp(context.Background(), nil, sa)
	if err != nil {
		log.Info("Error", "err1", err)
		return nil, nil
	}
	return app, nil
}

// RandSeq oihweiouh
func RandSeq() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, 20)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
