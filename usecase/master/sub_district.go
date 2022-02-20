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

type SubDistrictUsecaseItf interface {
	GetListSubDistrict(pagination gen.PaginationData, filter domain.SubDistrictFilter) ([]domain.SubDistrict, gen.PaginationData, string, error)
}

type SubDistrictUsecase struct {
	Repo   master.SubDistrictRepoItf
	DBList *infra.DatabaseList
	Log    *logrus.Logger
}

func newSubDistrictUsecase(r repo.Repo, logger *logrus.Logger, dbList *infra.DatabaseList) SubDistrictUsecase {
	return SubDistrictUsecase{
		Repo:   r.Master.SubDistrict,
		Log:    logger,
		DBList: dbList,
	}
}

func (sdu SubDistrictUsecase) GetListSubDistrict(pagination gen.PaginationData, filter domain.SubDistrictFilter) ([]domain.SubDistrict, gen.PaginationData, string, error) {
	data, err := sdu.Repo.GetListSubDistrict(pagination, filter)
	if err != nil {
		sdu.Log.WithField("filter", utils.StructToString(filter)).WithError(err).Error("GetListSubDistrict | fail to get sub district list from repo")
		return data, pagination, "", err
	}

	count, page, err := sdu.Repo.GetTotalDataSubDistrict(pagination, filter)
	if err != nil {
		sdu.Log.WithField("filter", utils.StructToString(filter)).WithError(err).Error("error get total data sub district from repo")
		return data, pagination, "", err
	}

	pagination.TotalData = int(count)
	pagination.TotalPage = int(page)

	return data, pagination, general.SourceFromDB, nil
}
