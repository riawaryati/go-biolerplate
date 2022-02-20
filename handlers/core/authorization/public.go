package authorization

import (
	"context"
	"crypto/sha256"
	"fmt"
	"net/http"
	"time"

	cg "github.com/furee/backend/constants/general"
	dg "github.com/furee/backend/domain/general"
	"github.com/furee/backend/handlers"
	"github.com/furee/backend/utils"
	"github.com/sirupsen/logrus"
)

type PublicHandler struct {
	log  *logrus.Logger
	Conf *dg.SectionService
}

func NewPublicHandler(conf *dg.SectionService, logger *logrus.Logger) PublicHandler {
	return PublicHandler{
		log:  logger,
		Conf: conf,
	}
}

type Session struct{}

func (ph PublicHandler) AuthValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		respData := handlers.ResponseData{
			Status: cg.Fail,
		}

		authorization := req.Header.Get("Authorization")
		authorizationID := req.Header.Get("Authorization-ID")

		if authorization == "" {
			respData.Message = "Token Not Valid"
			handlers.WriteResponse(res, respData, http.StatusUnauthorized)
			return
		}

		if authorizationID == "" {
			respData.Message = "Token Not Valid"
			handlers.WriteResponse(res, respData, http.StatusUnauthorized)
			return
		}

		authUnix, err := utils.StrToInt64(authorizationID)
		if err != nil {
			respData.Message = "Token Not Valid"
			handlers.WriteResponse(res, respData, http.StatusUnauthorized)
			return
		}

		authTime := time.Unix(authUnix, 0)
		if time.Now().UTC().Unix() > (authTime.UTC().Add(cg.Time1Min)).Unix() {
			respData.Message = "Token Not Valid"
			handlers.WriteResponse(res, respData, http.StatusUnauthorized)
			return
		}

		authCompareByte := sha256.Sum256([]byte(fmt.Sprintf("%s%s", ph.Conf.Authorization.Public.SecretKey, authorizationID)))
		authCompare := fmt.Sprintf("%x", authCompareByte)

		if authCompare != authorization {
			respData.Message = "Token Not Valid"
			handlers.WriteResponse(res, respData, http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(req.Context(), Session{}, authorization)
		req = req.WithContext(ctx)

		next.ServeHTTP(res, req)
	})
}
