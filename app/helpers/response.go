package helpers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	ResponseWritter http.ResponseWriter
}

type TemplateResponse struct {
	Code    int `json:"code"`
	Data    interface{}
	Message string
}

type TemplateResponseLogin struct {
	Code    int `json:"code"`
	Data    interface{}
	Token   string
	Message string
}

func (r *Response) SendOk(body interface{}, message string) {
	resp := &TemplateResponse{
		Code:    200,
		Data:    body,
		Message: message,
	}
	setHeader(r.ResponseWritter, 200)
	setJson(r.ResponseWritter, *resp)
}

func (r *Response) SendOkWithToken(body interface{}, token string, message string) {
	resp := &TemplateResponseLogin{
		Code:    200,
		Data:    body,
		Token:   "Bearer " + token,
		Message: message,
	}
	setHeader(r.ResponseWritter, 200)
	setJsonWithToken(r.ResponseWritter, *resp)
}

func (r *Response) SendError(body interface{}, message string) {
	resp := &TemplateResponse{
		Code:    500,
		Data:    nil,
		Message: message,
	}
	setHeader(r.ResponseWritter, 500)
	setJson(r.ResponseWritter, *resp)
}

func setHeader(rw http.ResponseWriter, status int) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status)
}
func setJson(rw http.ResponseWriter, resp TemplateResponse) {
	json.NewEncoder(rw).Encode(resp)
}
func setJsonWithToken(rw http.ResponseWriter, resp TemplateResponseLogin) {
	json.NewEncoder(rw).Encode(resp)
}
