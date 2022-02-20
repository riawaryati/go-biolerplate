package order

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	cg "github.com/furee/backend/constants/general"
	"github.com/furee/backend/domain/general"
	du "github.com/furee/backend/domain/order"
	"github.com/furee/backend/handlers"
	"github.com/furee/backend/usecase"
	uu "github.com/furee/backend/usecase/order"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gopkg.in/dealancer/validate.v2"
)

type OrderDataHandler struct {
	Usecase uu.OrderDataUsecaseItf
	conf    *general.SectionService
	log     *logrus.Logger
}

func newOrderHandler(uc usecase.Usecase, conf *general.SectionService, logger *logrus.Logger) OrderDataHandler {
	return OrderDataHandler{
		Usecase: uc.Order.Order,
		conf:    conf,
		log:     logger,
	}
}

func (ch OrderDataHandler) GetList(res http.ResponseWriter, req *http.Request) {
	respData := &handlers.ResponseData{
		Status: cg.Fail,
	}

	message := ""
	orders, err := ch.Usecase.GetList()
	if err != nil {
		// if orders == nil {
		// 	message = "No data found"
		// }
		message := err.Error()

		respData.Message = message
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	respData = &handlers.ResponseData{
		Status:  cg.Success,
		Message: message,
		Detail:  orders,
	}

	handlers.WriteResponse(res, respData, http.StatusOK)
}

func (ch OrderDataHandler) GetByID(res http.ResponseWriter, req *http.Request) {
	respData := &handlers.ResponseData{
		Status: cg.Fail,
	}

	message := ""
	orderidParam, ok := mux.Vars(req)["orderid"]

	if !ok {
		message = "Url Param 'orderid' is missing"
		respData.Message = message
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	orderid, err := strconv.ParseInt(orderidParam, 0, 64)
	if err != nil {
		message = "Invalid param order id"

		respData.Message = message
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	order, err := ch.Usecase.GetByID(orderid)

	if err != nil {
		// if order == nil {
		// 	message = "No data found"
		// } else {
		message = err.Error()
		// }

		respData.Message = message
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	if order == nil {
		message = "Data not found"

		respData.Message = message
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	// fmt.Print(order)
	respData = &handlers.ResponseData{
		Status:  cg.Success,
		Message: message,
		Detail:  order,
	}

	handlers.WriteResponse(res, respData, http.StatusOK)
}

func (ch OrderDataHandler) CreateOrder(res http.ResponseWriter, req *http.Request) {
	respData := &handlers.ResponseData{
		Status: cg.Fail,
	}

	var param du.OrderRequest

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		respData.Message = cg.HandlerErrorRequestDataEmpty
		handlers.WriteResponse(res, respData, http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(reqBody, &param)
	fmt.Println(param)
	if err != nil {
		respData.Message = cg.HandlerErrorRequestDataNotValid
		handlers.WriteResponse(res, respData, http.StatusBadRequest)
		return
	}

	err = validate.Validate(param)
	if err != nil {
		respData.Message = cg.HandlerErrorRequestDataFormatInvalid
		handlers.WriteResponse(res, respData, http.StatusBadRequest)
		return
	}

	message := ""
	orderId, err := ch.Usecase.CreateOrder(param)
	if err != nil {
		if orderId == 0 {
			message = "fail to create order"
		} else {
			message = err.Error()
		}

		respData.Message = message
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	respData = &handlers.ResponseData{
		Status:  cg.Success,
		Message: message,
	}

	handlers.WriteResponse(res, respData, http.StatusOK)
}

func (ch OrderDataHandler) UpdateOrder(res http.ResponseWriter, req *http.Request) {
	respData := &handlers.ResponseData{
		Status: cg.Fail,
	}

	message := ""
	orderidParam, ok := mux.Vars(req)["orderid"]

	if !ok {
		message = "Url Param 'orderid' is missing"
		respData.Message = message
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	orderid, err := strconv.ParseInt(orderidParam, 0, 64)
	if err != nil {
		message = "Invalid param order id"

		respData.Message = message
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	var param du.OrderRequest

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		respData.Message = cg.HandlerErrorRequestDataEmpty
		handlers.WriteResponse(res, respData, http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(reqBody, &param)
	if err != nil {
		respData.Message = cg.HandlerErrorRequestDataNotValid
		handlers.WriteResponse(res, respData, http.StatusBadRequest)
		return
	}

	param.OrderID = orderid
	err = validate.Validate(param)
	if err != nil {
		respData.Message = cg.HandlerErrorRequestDataFormatInvalid
		handlers.WriteResponse(res, respData, http.StatusBadRequest)
		return
	}

	updated, err := ch.Usecase.UpdateOrder(param)
	if err != nil {
		message = err.Error()

		respData.Message = message
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	respData = &handlers.ResponseData{
		Status:  cg.Success,
		Message: message,
		Detail:  updated,
	}

	handlers.WriteResponse(res, respData, http.StatusOK)
}

func (ch OrderDataHandler) DeleteByID(res http.ResponseWriter, req *http.Request) {
	respData := &handlers.ResponseData{
		Status: cg.Fail,
	}

	message := ""
	orderidParam, ok := mux.Vars(req)["orderid"]
	fmt.Print(orderidParam)
	if !ok {
		message = "Url Param 'orderid' is missing"
		respData.Message = message
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	orderid, err := strconv.ParseInt(orderidParam, 0, 64)
	if err != nil {
		message = "Invalid param order id"

		respData.Message = message
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	deleted, err := ch.Usecase.DeleteByID(orderid)

	if err != nil {
		message = err.Error()

		respData.Message = message
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	respData = &handlers.ResponseData{
		Status:  cg.Success,
		Message: message,
		Detail:  deleted,
	}

	handlers.WriteResponse(res, respData, http.StatusOK)
}
