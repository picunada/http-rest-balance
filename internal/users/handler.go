package users

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/picunada/http-rest-balance/internal/handlers"
	"github.com/picunada/http-rest-balance/pkg/logging"
	"net/http"
)

type handler struct {
	logger logging.Logger
}

func NewHandler(logger logging.Logger) hanlders.Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET("/users", h.GetAllUsers)
	router.GET("/users/:uuid", h.GetUserById)
	router.PUT("/users/:uuid", h.UpdateUser)
	router.POST("/users", h.CreateUser)
}

func (h *handler) GetAllUsers(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte(fmt.Sprintf("this is balance")))
}

func (h *handler) GetUserById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte(fmt.Sprintf("this is balance")))
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte(fmt.Sprintf("this is balance")))
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte(fmt.Sprintf("this is balance")))
}
