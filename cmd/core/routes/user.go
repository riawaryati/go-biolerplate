package routes

import (
	"net/http"

	"github.com/furee/backend/domain/general"
	"github.com/furee/backend/handlers/core"
	"github.com/gorilla/mux"
)

func getUser(router, routerJWT *mux.Router, conf *general.SectionService, handler core.Handler) {
	router.HandleFunc("/verify-user", handler.User.User.VerifyOTP).Methods(http.MethodPost)
	router.HandleFunc("/login", handler.User.User.LoginUser).Methods(http.MethodPost)
}
