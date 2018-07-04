package sewan_go_sdk

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func HttpResponseFake_OK_json() *http.Response {
	response := http.Response{}
	response.Status = "200 OK"
	response.StatusCode = http.StatusOK
	response.Header = map[string][]string{}
	response.Header.Add("Content-Type", "application/json")
	Resp_Body_json, _ := json.Marshal(Resp_Body{Detail: "a simple json"})
	response.Body = ioutil.NopCloser(bytes.NewBuffer(Resp_Body_json))
	return &response
}

func HttpResponseFake_OK_template_list_json() *http.Response {
	response := http.Response{}
	response.Status = "200 OK"
	response.StatusCode = http.StatusOK
	response.Header = map[string][]string{}
	response.Header.Add("Content-Type", "application/json")
	Resp_Body_json, _ := json.Marshal(TEMPLATES_LIST)
	response.Body = ioutil.NopCloser(bytes.NewBuffer(Resp_Body_json))
	return &response
}

func HttpResponseFake_500_texthtml() *http.Response {
	response := http.Response{}
	response.Status = "500 Internal Server Error"
	response.StatusCode = http.StatusInternalServerError
	response.Header = map[string][]string{}
	response.Header.Add("Content-Type", "text/html")
	response.Body = ioutil.NopCloser(bytes.NewBufferString("<h1>Server Error (500)</h1>"))
	return &response
}

func HttpResponseFake_500_json() *http.Response {
	response := http.Response{}
	response.Status = "500 Internal Server Error"
	response.StatusCode = http.StatusInternalServerError
	response.Header = map[string][]string{}
	response.Header.Add("Content-Type", "application/json")
	Resp_Body_json, _ := json.Marshal(Resp_Body{Detail: "a json response Resp_Body"})
	response.Body = ioutil.NopCloser(bytes.NewBuffer(Resp_Body_json))
	return &response
}

func HttpResponseFake_OK_txthtml() *http.Response {
	response := http.Response{}
	response.Status = "200 OK"
	response.StatusCode = http.StatusOK
	response.Header = map[string][]string{}
	response.Header.Add("Content-Type", "text/html")
	response.Body = ioutil.NopCloser(bytes.NewBufferString("<h1>An html text</h1>"))
	return &response
}

func HttpResponseFake_OK_no_content() *http.Response {
	response := http.Response{}
	response.Status = "200 OK"
	response.StatusCode = http.StatusOK
	return &response
}

func HttpResponseFake_OK_wrongjson() *http.Response {
	response := http.Response{}
	response.Status = "200 OK"
	response.StatusCode = http.StatusOK
	response.Header = map[string][]string{}
	response.Header.Add("Content-Type", "application/json")
	response.Body = ioutil.NopCloser(bytes.NewBufferString("a bad formated json"))
	return &response
}

func HttpResponseFake_OK_image() *http.Response {
	response := http.Response{}
	response.Status = "200 OK"
	response.StatusCode = http.StatusOK
	response.Header = map[string][]string{}
	response.Header.Add("Content-Type", "image")
	return &response
}
