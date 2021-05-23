package lib

import (
	"context"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func FirebaseApp(ctx context.Context) (app *firebase.App) {
	opt := option.WithCredentialsFile("service-account.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return
	}
	return app
}
