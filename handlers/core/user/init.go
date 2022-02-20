package user

import (
	"github.com/furee/backend/domain/general"
	"github.com/furee/backend/usecase"
	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	User UserDataHandler
}

func NewHandler(uc usecase.Usecase, conf *general.SectionService, logger *logrus.Logger) UserHandler {
	return UserHandler{
		User: newUserHandler(uc, conf, logger),
	}
}
