package configs

import (
	"context"
	"encoding/json"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type FirebaseCredentials struct {
	Type                    string `json:"type"`
	ProjectID               string `json:"project_id"`
	PrivateKeyID            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientID                string `json:"client_id"`
	AuthURI                 string `json:"auth_uri"`
	TokenURI                string `json:"token_uri"`
	AuthProviderX509CertURL string `json:"auth_provider_x509_cert_url"`
	ClientX509CertURL       string `json:"client_x509_cert_url"`
	UniverseDomain          string `json:"universe_domain"`
}

func InitFirebaseApp(ctx context.Context, credentialsJSONString string) (*firebase.App, error) {
	credentialsJSON := []byte(credentialsJSONString)

	var credentials FirebaseCredentials
	err := json.Unmarshal(credentialsJSON, &credentials)
	if err != nil {
		return nil, err
	}

	config := &firebase.Config{ProjectID: FirebaseProjectId}
	opt := option.WithCredentialsJSON(credentialsJSON)

	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		return nil, err
	}

	return app, nil
}
