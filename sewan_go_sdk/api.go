package sewan_go_sdk

import(
  "net/http"
)

// API is the interface used to communicate with the  API
type API struct {
  Token string
  URL string
  Client *http.Client
}

// New creates a ready-to-use SDK client
func New(token string, url string) (*API, error) {
  api := &API{
    Token: token,
    URL: url,
    Client: &http.Client{},
  }
  return api, nil
}
