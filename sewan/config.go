package sewan

import(
  "net/http"
  sdk "terraform-provider-sewan/sewan_go_sdk"
)

// Config contains scaleway configuration values
type Config struct{
  Api_token string
  Api_url string
}

type API struct {
  Token string
  URL string
  Client *http.Client
}

// Client contains scaleway api clients
type Client struct {
  sewan *sdk.API
}

func (c *Config) Client() (*Client, error) {
  api, err := sdk.New(
    c.Api_token,
    c.Api_url,
  )
  if err != nil {
    return nil, err
  }

  // TODO : validate token here

  return &Client{api}, nil
}
