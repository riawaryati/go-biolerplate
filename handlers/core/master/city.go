package master

import (
	"net/http"
	"strconv"

	cg "github.com/furee/backend/constants/general"
	"github.com/furee/backend/domain/general"
	dm "github.com/furee/backend/domain/master"
	"github.com/furee/backend/handlers"
	"github.com/furee/backend/usecase"
	um "github.com/furee/backend/usecase/master"
	"github.com/furee/backend/utils"
	"github.com/sirupsen/logrus"
	"gopkg.in/guregu/null.v4"
)

type CityHandler struct {
	Usecase um.CityUsecaseItf
	conf    *general.SectionService
	log     *logrus.Logger
}

func newCityHandler(uc usecase.Usecase, conf *general.SectionService, logger *logrus.Logger) CityHandler {
	return CityHandler{
		Usecase: uc.Master.City,
		conf:    conf,
		log:     logger,
	}
}

func (ch CityHandler) GetListCity(res http.ResponseWriter, req *http.Request) {
	respData := &handlers.ResponseData{
		Status: cg.Fail,
	}

	var tableFilter dm.CityFilter
	var err error

	paginationData := general.GetPagination()

	// Check province name value
	if req.FormValue("name") != "" {
		tableFilter.Name = null.StringFrom(req.FormValue("name"))
	}

	// Check country_id value
	if req.FormValue("province-id") != "" {
		provinceID, err := utils.StrToInt64(req.FormValue("province-id"))
		if err != nil {
			respData.Message = cg.HandlerErrorRequestDataFormatInvalid
			handlers.WriteResponse(res, respData, http.StatusBadRequest)
			return
		}

		tableFilter.ProvinceID = null.IntFrom(provinceID)
	}

	// Check sort value
	if req.FormValue("sort") != "" {
		paginationData.Sort = req.FormValue(("sort"))
	}

	// Check page value. If exist, convert to int
	if req.FormValue("page") != "" {
		paginationData.Page, err = strconv.Atoi(req.FormValue("page"))
		if err != nil {
			respData.Message = cg.HandlerErrorRequestDataFormatInvalid
			handlers.WriteResponse(res, respData, http.StatusBadRequest)
			return
		}
	}

	// Check orderby value.
	paginationData.OrderBy = null.StringFrom("city_id")
	if req.FormValue("order-by") != "" {
		paginationData.OrderBy.String = req.FormValue("order-by")
	}

	// Check isGetAll value.
	if req.FormValue("is-get-all") != "" {
		paginationData.IsGetAll = utils.GetBool(req.FormValue("is-get-all"))
	}

	// Check limit value. If exists, convert to int
	if req.FormValue("limit") != "" {
		paginationData.Limit, err = strconv.Atoi(req.FormValue("limit"))
		if err != nil {
			respData.Message = cg.HandlerErrorRequestDataFormatInvalid
			handlers.WriteResponse(res, respData, http.StatusBadRequest)
			return
		}
	}

	// Convert page to offset
	paginationData.SetOffset()

	data, paginationData, _, err := ch.Usecase.GetListCity(paginationData, tableFilter)
	if err != nil {
		respData.Message = "fail to get list city"
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	respData = &handlers.ResponseData{
		Status:  cg.Success,
		Message: "success get list city",
		Detail: general.ResponseData{
			Data:       data,
			Pagination: paginationData,
		},
	}

	handlers.WriteResponse(res, respData, http.StatusOK)
	return
}
