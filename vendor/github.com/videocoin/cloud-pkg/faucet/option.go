package faucet

import (
	"context"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Option func(*Client)

func WithBasicAuth(user, password string) Option {
	return func(c *Client) {
		c.user = user
		c.password = password
	}
}

func WithTokenSource(oauthClientID, accessToken string) Option {
	return func(c *Client) {
		cfg, err := google.JWTConfigFromJSON([]byte(accessToken))
		if err != nil {
			return
		}

		cfg.UseIDToken = true
		cfg.PrivateClaims = map[string]interface{}{
			"target_audience": oauthClientID,
		}

		c.cli = oauth2.NewClient(context.Background(), cfg.TokenSource(context.Background()))
	}
}
