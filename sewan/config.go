package sewan

import (
	"github.com/hashicorp/terraform/helper/schema"
)

type AirDrumClient struct {
	ApiToken   string
	Endpoint   string
	Timeout    int
	MaxRetries int
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	SewanClient := AirDrumClient{
		ApiToken:   d.Get("api_key").(string),
		Endpoint:   d.Get("endpoint").(string),
		Timeout:    d.Get("timeout").(int),
		MaxRetries: d.Get("max_retries").(int),
  }

	return &SewanClient, nil
}
