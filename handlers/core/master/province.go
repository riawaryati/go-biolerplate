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

type ProvinceHandler struct {
	Usecase um.ProvinceUsecaseItf
	conf    *general.SectionService
	log     *logrus.Logger
}

func newProvinceHandler(uc usecase.Usecase, conf *general.SectionService, logger *logrus.Logger) ProvinceHandler {
	return ProvinceHandler{
		Usecase: uc.Master.Province,
		conf:    conf,
		log:     logger,
	}
}

func (ph ProvinceHandler) GetListProvince(res http.ResponseWriter, req *http.Request) {
	respData := &handlers.ResponseData{
		Status: cg.Fail,
	}

	var tableFilter dm.ProvinceFilter
	var err error

	paginationData := general.GetPagination()

	// Check province name value
	if req.FormValue("name") != "" {
		tableFilter.Name = null.StringFrom(req.FormValue("name"))
	}

	// Check country_id value
	if req.FormValue("country-id") != "" {
		countryID, err := utils.StrToInt64(req.FormValue("country-id"))
		if err != nil {
			respData.Message = cg.HandlerErrorRequestDataFormatInvalid
			handlers.WriteResponse(res, respData, http.StatusBadRequest)
			return
		}

		tableFilter.CountryID = null.IntFrom(countryID)
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
	paginationData.OrderBy = null.StringFrom("province_id")
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

	data, paginationData, _, err := ph.Usecase.GetListProvince(paginationData, tableFilter)
	if err != nil {
		respData.Message = "fail to get list province"
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	respData = &handlers.ResponseData{
		Status:  cg.Success,
		Message: "success get list province",
		Detail: general.ResponseData{
			Data:       data,
			Pagination: paginationData,
		},
	}

	handlers.WriteResponse(res, respData, http.StatusOK)
	return
}
