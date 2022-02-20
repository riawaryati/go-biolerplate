package handlers

import (
	"encoding/json"
	"net/http"

	constants "github.com/furee/backend/constants/general"
	"github.com/furee/backend/domain/general"
)

type ResponseHTTP struct {
	StatusCode int
	Response   ResponseData
}

type ResponseData struct {
	Status  string      `json:"status"`
	Source  string      `json:"source,omitempty"`
	Message string      `json:"message,omitempty"`
	Detail  interface{} `json:"detail,omitempty"`
}

// Response is the new type for define all of the response from service
type Response interface{}

var (
	ErrRespServiceMaintance = ResponseHTTP{
		StatusCode: http.StatusServiceUnavailable,
		Response:   ResponseData{Status: constants.Fail}}
	ErrRespUnauthorize = ResponseHTTP{
		StatusCode: http.StatusUnauthorized,
		Response:   ResponseData{Status: constants.Fail}}
	ErrRespAuthInvalid = ResponseHTTP{
		StatusCode: http.StatusUnauthorized,
		Response:   ResponseData{Status: constants.Fail}}
	ErrRespBadRequest = ResponseHTTP{
		StatusCode: http.StatusBadRequest,
		Response:   ResponseData{Status: constants.Fail}}
	ErrRespInternalServer = ResponseHTTP{
		StatusCode: http.StatusServiceUnavailable,
		Response:   ResponseData{Status: constants.Fail}}
)

func WriteResponse(res http.ResponseWriter, resp Response, code int) {
	res.Header().Set("Content-Type", "application/json")
	r, _ := json.Marshal(resp)

	res.WriteHeader(code)
	res.Write(r)
	return
}

type Error struct {
	Id     string `json:"id"`
	Status string `json:"status"`
	Title  string `json:"title"`
}

func NewError(id string, status string, title string) *Error {
	return &Error{
		Id:     id,
		Status: status,
		Title:  title,
	}
}

func (rd *ResponseData) GenerateErrorResponse(data *general.ResponseData, errorMsg string) {
	data.Error = errorMsg
	rd.Detail = data
}
