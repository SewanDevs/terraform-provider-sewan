package sewan

import (
	"net/http"
	sdk "terraform-provider-sewan/sewan_go_sdk"
)

type Config struct {
	Api_token string
	Api_url   string
}

type API struct {
	Token  string
	URL    string
	Client *http.Client
}

type Client struct {
	sewan *sdk.API
}

func (c *Config) Client() (*Client, error) {
	apiTooler := sdk.APITooler{
		Api: sdk.AirDrumAPIer{},
	}
	api := apiTooler.New(
		c.Api_token,
		c.Api_url,
	)
	err := apiTooler.CheckStatus(api)
	return &Client{api}, err
}
