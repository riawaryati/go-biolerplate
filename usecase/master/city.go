package master

import (
	"github.com/furee/backend/constants/general"
	gen "github.com/furee/backend/domain/general"
	domain "github.com/furee/backend/domain/master"
	"github.com/furee/backend/infra"
	"github.com/furee/backend/repo"
	"github.com/furee/backend/repo/master"
	"github.com/furee/backend/utils"
	"github.com/sirupsen/logrus"
)

type CityUsecaseItf interface {
	GetListCity(pagination gen.PaginationData, filter domain.CityFilter) ([]domain.City, gen.PaginationData, string, error)
}

type CityUsecase struct {
	Repo   master.CityRepoItf
	DBList *infra.DatabaseList
	Log    *logrus.Logger
}

func newCityUsecase(r repo.Repo, logger *logrus.Logger, dbList *infra.DatabaseList) CityUsecase {
	return CityUsecase{
		Repo:   r.Master.City,
		Log:    logger,
		DBList: dbList,
	}
}

func (cu CityUsecase) GetListCity(pagination gen.PaginationData, filter domain.CityFilter) ([]domain.City, gen.PaginationData, string, error) {
	data, err := cu.Repo.GetListCity(pagination, filter)
	if err != nil {
		cu.Log.WithField("filter", utils.StructToString(filter)).WithError(err).Error("GetListCity | fail to get city list from repo")
		return data, pagination, "", err
	}

	count, page, err := cu.Repo.GetTotalDataCity(pagination, filter)
	if err != nil {
		cu.Log.WithField("filter", utils.StructToString(filter)).WithError(err).Error("error get total data city from repo")
		return data, pagination, "", err
	}

	pagination.TotalData = int(count)
	pagination.TotalPage = int(page)

	return data, pagination, general.SourceFromDB, nil
}
