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

type CountryUsecaseItf interface {
	GetListCountry(pagination gen.PaginationData, filter domain.CountryFilter) ([]domain.Country, gen.PaginationData, string, error)
}

type CountryUsecase struct {
	Repo   master.CountryRepoItf
	DBList *infra.DatabaseList
	Log    *logrus.Logger
}

func newCountryUsecase(r repo.Repo, logger *logrus.Logger, dbList *infra.DatabaseList) CountryUsecase {
	return CountryUsecase{
		Repo:   r.Master.Country,
		Log:    logger,
		DBList: dbList,
	}
}

func (cu CountryUsecase) GetListCountry(pagination gen.PaginationData, filter domain.CountryFilter) ([]domain.Country, gen.PaginationData, string, error) {
	data, err := cu.Repo.GetListCountry(pagination, filter)
	if err != nil {
		cu.Log.WithField("filter", utils.StructToString(filter)).WithError(err).Error("GetListCountry | fail to get country list from repo")
		return data, pagination, "", err
	}

	count, page, err := cu.Repo.GetTotalDataCountry(pagination, filter)
	if err != nil {
		cu.Log.WithField("filter", utils.StructToString(filter)).WithError(err).Error("error get total data country from repo")
		return data, pagination, "", err
	}

	pagination.TotalData = int(count)
	pagination.TotalPage = int(page)

	return data, pagination, general.SourceFromDB, nil
}
