package user

import (
	"github.com/furee/backend/domain/general"
	"github.com/furee/backend/infra"
	"github.com/furee/backend/repo"
	"github.com/sirupsen/logrus"
)

type UserUsecase struct {
	User UserDataUsecaseItf
}

func NewUsecase(repo repo.Repo, conf *general.SectionService, dbList *infra.DatabaseList, logger *logrus.Logger) UserUsecase {
	return UserUsecase{
		User: newUserDataUsecase(repo, conf, logger, dbList),
	}
}
