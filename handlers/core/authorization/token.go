package authorization

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	cg "github.com/furee/backend/constants/general"
	dg "github.com/furee/backend/domain/general"
	"github.com/furee/backend/handlers"
	"github.com/furee/backend/utils"
	"github.com/sirupsen/logrus"
)

type TokenHandler struct {
	log  *logrus.Logger
	Conf *dg.SectionService
}

func NewTokenHandler(conf *dg.SectionService, logger *logrus.Logger) TokenHandler {
	utils.InitJWTConfig(conf.Authorization.JWT)
	return TokenHandler{
		log:  logger,
		Conf: conf,
	}
}

func (th TokenHandler) JWTValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		respData := handlers.ResponseData{
			Status: cg.Fail,
		}

		//List of URL that bypass this JWTValidator middleware
		if req.URL.Path == "/api/v1/renew-token" {
			next.ServeHTTP(res, req)
			return
		}

		authorizationHeader := req.Header.Get("Authorization")
		if !strings.Contains(authorizationHeader, "Bearer") {
			th.log.Error(fmt.Errorf("Invalid Token Format"))
			respData.Message = "Invalid Token Format"
			handlers.WriteResponse(res, respData, http.StatusBadRequest)
			return
		}
		accessToken := strings.Replace(authorizationHeader, "Bearer ", "", -1)

		claims, err := utils.CheckAccessToken(accessToken)
		if err != nil {
			respData.Message = "Token expired"
			handlers.WriteResponse(res, respData, http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(req.Context(), "session", claims["session"])
		req = req.WithContext(ctx)

		next.ServeHTTP(res, req)
	})
}

func (th TokenHandler) RenewAccessToken(res http.ResponseWriter, req *http.Request) {
	respData := &handlers.ResponseData{
		Status: cg.Fail,
	}

	authorizationHeader := req.Header.Get("Authorization")
	if !strings.Contains(authorizationHeader, "Bearer") {
		th.log.WithField("renew access token", authorizationHeader).Error("Invalid Authorization Format")
		respData.Message = "Invalid Authorization Format"
		handlers.WriteResponse(res, respData, http.StatusBadRequest)
		return
	}

	accessToken := strings.Replace(authorizationHeader, "Bearer ", "", -1)

	token, err := utils.RenewAccessToken(accessToken)
	if err != nil {
		th.log.WithField("renew access token", accessToken).WithError(err).Error("Error Renew Token")
		respData.Message = "Fail to Renew Token"
		handlers.WriteResponse(res, respData, http.StatusBadRequest)
		return
	}

	tokenExpired := time.Now().UTC().Add(time.Duration(th.Conf.Authorization.JWT.AccessTokenDuration) * time.Minute)

	respData = &handlers.ResponseData{
		Status:  cg.Success,
		Message: "success generate new access token",
		Detail: dg.RenewToken{
			Token:        token,
			TokenExpired: tokenExpired.Format(time.RFC3339),
		},
	}

	handlers.WriteResponse(res, respData, http.StatusOK)
	return
}
