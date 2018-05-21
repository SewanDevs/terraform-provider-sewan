package sewan_go_sdk

import(
  "net/http"
)

type Config struct{
  Api_token string `json:"api_token"`
  Api_url string `json:"api_url"`
}

type HTTP_Clienter interface{
  GetConf()
  Get_HTTP_client()
}

type Http_client struct{
  Conf Config
  Net_http_client *http.Client
}

func (cl Http_client) GetConf() Config {
    return cl.Conf
}

func (cl Http_client) Get_HTTP_client() *http.Client {
    return cl.Net_http_client
}

func CreateClient(c Config) Http_client {
  client := Http_client{c, &http.Client{}}
  return client
}
