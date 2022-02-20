package routes

import (
	"net/http"

	"github.com/furee/backend/domain/general"
	"github.com/furee/backend/handlers/core"
	"github.com/gorilla/mux"
)

func getMasterData(router, routerJWT *mux.Router, conf *general.SectionService, handler core.Handler) {
	routerJWT.HandleFunc("/country", handler.Master.Country.GetListCountry).Methods(http.MethodGet)
	routerJWT.HandleFunc("/province", handler.Master.Province.GetListProvince).Methods(http.MethodGet)
	routerJWT.HandleFunc("/city", handler.Master.City.GetListCity).Methods(http.MethodGet)
	routerJWT.HandleFunc("/district", handler.Master.District.GetListDistrict).Methods(http.MethodGet)
	routerJWT.HandleFunc("/sub-district", handler.Master.SubDistrict.GetListSubDistrict).Methods(http.MethodGet)
}
