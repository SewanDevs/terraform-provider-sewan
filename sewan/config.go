package sewan

/*import (
	"github.com/hashicorp/terraform/helper/schema"
)*/

type AirDrumConfig struct {
	AirDrumToken string
  AirDrumURL string
  Timeout int
  Max_retries int
}
