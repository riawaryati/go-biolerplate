package order

import (
	"github.com/furee/backend/domain/general"
	"github.com/furee/backend/infra"
	"github.com/furee/backend/repo"
	"github.com/sirupsen/logrus"
)

type OrderUsecase struct {
	Order OrderDataUsecaseItf
}

func NewUsecase(repo repo.Repo, conf *general.SectionService, dbList *infra.DatabaseList, logger *logrus.Logger) OrderUsecase {
	return OrderUsecase{
		Order: newOrderDataUsecase(repo, conf, logger, dbList),
	}
}
