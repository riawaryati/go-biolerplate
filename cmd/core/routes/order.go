package routes

import (
	"net/http"

	"github.com/furee/backend/domain/general"
	"github.com/furee/backend/handlers/core"
	"github.com/gorilla/mux"
)

func getOrder(router, routerJWT *mux.Router, conf *general.SectionService, handler core.Handler) {
	router.HandleFunc("/orders/{orderid}", handler.Order.Order.UpdateOrder).Methods(http.MethodPut)
	router.HandleFunc("/orders", handler.Order.Order.CreateOrder).Methods(http.MethodPost)
	router.HandleFunc("/orders", handler.Order.Order.GetList).Methods(http.MethodGet)
	router.HandleFunc("/orders/{orderid}", handler.Order.Order.DeleteByID).Methods(http.MethodDelete)
	router.HandleFunc("/orders/{orderid}", handler.Order.Order.GetByID).Methods(http.MethodGet)
}
