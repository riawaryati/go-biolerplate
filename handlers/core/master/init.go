package master

import (
	"github.com/furee/backend/domain/general"
	"github.com/furee/backend/usecase"
	"github.com/sirupsen/logrus"
)

type MasterHandler struct {
	Country     CountryHandler
	Province    ProvinceHandler
	City        CityHandler
	District    DistrictHandler
	SubDistrict SubDistrictHandler
}

func NewHandler(uc usecase.Usecase, conf *general.SectionService, logger *logrus.Logger) MasterHandler {
	return MasterHandler{
		Country:     newCountryHandler(uc, conf, logger),
		Province:    newProvinceHandler(uc, conf, logger),
		City:        newCityHandler(uc, conf, logger),
		District:    newDistrictHandler(uc, conf, logger),
		SubDistrict: newSubDistrictHandler(uc, conf, logger),
	}
}
