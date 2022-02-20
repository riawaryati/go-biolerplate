package core

import (
	"github.com/furee/backend/domain/general"
	"github.com/furee/backend/handlers/core/authorization"
	"github.com/furee/backend/handlers/core/master"
	"github.com/furee/backend/handlers/core/order"
	"github.com/furee/backend/handlers/core/user"
	"github.com/furee/backend/usecase"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	Token  authorization.TokenHandler
	Public authorization.PublicHandler
	Master master.MasterHandler
	User   user.UserHandler
	Order  order.OrderHandler
}

func NewHandler(uc usecase.Usecase, conf *general.SectionService, logger *logrus.Logger) Handler {
	return Handler{
		Token:  authorization.NewTokenHandler(conf, logger),
		Public: authorization.NewPublicHandler(conf, logger),
		Master: master.NewHandler(uc, conf, logger),
		User:   user.NewHandler(uc, conf, logger),
		Order:  order.NewHandler(uc, conf, logger),
	}
}
