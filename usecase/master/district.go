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

type DistrictUsecaseItf interface {
	GetListDistrict(pagination gen.PaginationData, filter domain.DistrictFilter) ([]domain.District, gen.PaginationData, string, error)
}

type DistrictUsecase struct {
	Repo   master.DistrictRepoItf
	DBList *infra.DatabaseList
	Log    *logrus.Logger
}

func newDistrictUsecase(r repo.Repo, logger *logrus.Logger, dbList *infra.DatabaseList) DistrictUsecase {
	return DistrictUsecase{
		Repo:   r.Master.District,
		Log:    logger,
		DBList: dbList,
	}
}

func (du DistrictUsecase) GetListDistrict(pagination gen.PaginationData, filter domain.DistrictFilter) ([]domain.District, gen.PaginationData, string, error) {
	data, err := du.Repo.GetListDistrict(pagination, filter)
	if err != nil {
		du.Log.WithField("filter", utils.StructToString(filter)).WithError(err).Error("GetListDistrict | fail to get district list from repo")
		return data, pagination, "", err
	}

	count, page, err := du.Repo.GetTotalDataDistrict(pagination, filter)
	if err != nil {
		du.Log.WithField("filter", utils.StructToString(filter)).WithError(err).Error("error get total data district from repo")
		return data, pagination, "", err
	}

	pagination.TotalData = int(count)
	pagination.TotalPage = int(page)

	return data, pagination, general.SourceFromDB, nil
}
