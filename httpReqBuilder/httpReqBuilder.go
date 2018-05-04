package httpReqBuilder

import (
	//"fmt"
	//"net/http"
)

type createReq struct{
  token string
}

func NewCreateReq(token string) createReq{
  request := createReq{token}
  return request
}

func Create() error {

	return nil
}

func Get() error {
	return nil
}

func Put() error {
	return nil
}

func Delete() error {
	return nil
}
