package usecase

import (
	"github.com/furee/backend/domain/general"
	"github.com/furee/backend/infra"
	"github.com/furee/backend/repo"
	"github.com/furee/backend/usecase/master"
	"github.com/furee/backend/usecase/order"
	"github.com/furee/backend/usecase/user"
	"github.com/sirupsen/logrus"
)

type Usecase struct {
	Master master.MasterUsecase
	User   user.UserUsecase
	Order  order.OrderUsecase
}

func NewUsecase(repo repo.Repo, conf *general.SectionService, dbList *infra.DatabaseList, logger *logrus.Logger) Usecase {
	return Usecase{
		Master: master.NewUsecase(repo, conf, dbList, logger),
		User:   user.NewUsecase(repo, conf, dbList, logger),
		Order:  order.NewUsecase(repo, conf, dbList, logger),
	}
}
