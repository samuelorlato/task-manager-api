package configs

import (
	"context"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func InitFirebaseApp(ctx context.Context) (*firebase.App, error) {
	config := &firebase.Config{ProjectID: "task-manager-3d540"}
	opt := option.WithCredentialsFile("../../internal/configs/credentials.json")
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		return nil, err
	}

	return app, nil
}
