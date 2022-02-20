package user

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	cg "github.com/furee/backend/constants/general"
	"github.com/furee/backend/domain/general"
	du "github.com/furee/backend/domain/user"
	"github.com/furee/backend/handlers"
	"github.com/furee/backend/usecase"
	uu "github.com/furee/backend/usecase/user"
	"github.com/sirupsen/logrus"
	"gopkg.in/dealancer/validate.v2"
)

type UserDataHandler struct {
	Usecase uu.UserDataUsecaseItf
	conf    *general.SectionService
	log     *logrus.Logger
}

func newUserHandler(uc usecase.Usecase, conf *general.SectionService, logger *logrus.Logger) UserDataHandler {
	return UserDataHandler{
		Usecase: uc.User.User,
		conf:    conf,
		log:     logger,
	}
}

func (ch UserDataHandler) VerifyOTP(res http.ResponseWriter, req *http.Request) {
	respData := &handlers.ResponseData{
		Status: cg.Fail,
	}

	var param du.VerifyOTPRequest

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		respData.Message = cg.HandlerErrorRequestDataEmpty
		handlers.WriteResponse(res, respData, http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(reqBody, &param)
	if err != nil {
		respData.Message = cg.HandlerErrorRequestDataNotValid
		handlers.WriteResponse(res, respData, http.StatusBadRequest)
		return
	}

	err = validate.Validate(param)
	if err != nil {
		respData.Message = cg.HandlerErrorRequestDataFormatInvalid
		handlers.WriteResponse(res, respData, http.StatusBadRequest)
		return
	}

	jwt, message, err := ch.Usecase.VerifyOTP(param)
	if err != nil {
		if message == "" {
			message = "fail to verify otp user"
		}

		respData.Message = message
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	respData = &handlers.ResponseData{
		Status:  cg.Success,
		Message: message,
		Detail:  jwt,
	}

	handlers.WriteResponse(res, respData, http.StatusOK)
	return
}

func (ch UserDataHandler) LoginUser(res http.ResponseWriter, req *http.Request) {
	respData := &handlers.ResponseData{
		Status: cg.Fail,
	}

	var param du.UserLoginRequest

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		respData.Message = cg.HandlerErrorRequestDataEmpty
		handlers.WriteResponse(res, respData, http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(reqBody, &param)
	if err != nil {
		respData.Message = cg.HandlerErrorRequestDataNotValid
		handlers.WriteResponse(res, respData, http.StatusBadRequest)
		return
	}

	err = validate.Validate(param)
	if err != nil {
		respData.Message = cg.HandlerErrorRequestDataFormatInvalid
		handlers.WriteResponse(res, respData, http.StatusBadRequest)
		return
	}

	message, err := ch.Usecase.LoginUser(param)
	if err != nil {
		if message == "" {
			message = "fail to send otp user"
		}

		respData.Message = message
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	respData = &handlers.ResponseData{
		Status:  cg.Success,
		Message: message,
	}

	handlers.WriteResponse(res, respData, http.StatusOK)
	return
}
