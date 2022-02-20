package order

import (
	"github.com/furee/backend/domain/general"
	"github.com/furee/backend/usecase"
	"github.com/sirupsen/logrus"
)

type OrderHandler struct {
	Order OrderDataHandler
}

func NewHandler(uc usecase.Usecase, conf *general.SectionService, logger *logrus.Logger) OrderHandler {
	return OrderHandler{
		Order: newOrderHandler(uc, conf, logger),
	}
}
