package master

import (
	"github.com/furee/backend/infra"
	"github.com/sirupsen/logrus"
)

type MasterRepo struct {
	Country     CountryRepoItf
	SubDistrict SubDistrictRepoItf
	District    DistrictRepoItf
	City        CityRepoItf
	Province    ProvinceRepoItf
}

func NewMasterRepo(db *infra.DatabaseList, logger *logrus.Logger) MasterRepo {
	return MasterRepo{
		Country:     newCountryRepo(db),
		SubDistrict: newSubDistrictRepo(db),
		District:    newDistrictRepo(db),
		City:        newCityRepo(db),
		Province:    newProvinceRepo(db),
	}
}
